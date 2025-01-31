package jsonschema

import (
	"github.com/go-faster/errors"
)

// ReferenceResolver resolves JSON schema references.
type ReferenceResolver interface {
	ResolveReference(ref string) (*RawSchema, error)
}

// Parser parses JSON schemas.
type Parser struct {
	resolver   ReferenceResolver
	refcache   map[string]*Schema
	inferTypes bool
}

func NewParser(s Settings) *Parser {
	s.setDefaults()
	return &Parser{
		resolver:   s.Resolver,
		refcache:   map[string]*Schema{},
		inferTypes: s.InferTypes,
	}
}

func (p *Parser) Parse(schema *RawSchema) (*Schema, error) {
	return p.parse(schema, func(s *Schema) *Schema {
		return p.extendInfo(schema, s)
	})
}

func (p *Parser) parse(schema *RawSchema, hook func(*Schema) *Schema) (*Schema, error) {
	if schema == nil {
		return nil, nil
	}

	if ref := schema.Ref; ref != "" {
		s, err := p.resolve(ref)
		if err != nil {
			return nil, errors.Wrapf(err, "reference %q", ref)
		}
		return s, nil
	}

	switch {
	case len(schema.Enum) > 0:
		if p.inferTypes && schema.Type == "" {
			typ, err := inferEnumType(schema.Enum[0])
			if err != nil {
				return nil, errors.Wrap(err, "infer enum type")
			}
			schema.Type = typ
		}
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		values, err := parseEnumValues(SchemaType(schema.Type), schema.Enum)
		if err != nil {
			return nil, errors.Wrap(err, "parse enum")
		}

		return hook(&Schema{
			Type:   SchemaType(schema.Type),
			Format: Format(schema.Format),
			Enum:   values,
		}), nil
	case len(schema.OneOf) > 0:
		schemas, err := p.parseMany(schema.OneOf)
		if err != nil {
			return nil, errors.Wrapf(err, "oneOf")
		}

		return hook(&Schema{OneOf: schemas}), nil
	case len(schema.AnyOf) > 0:
		schemas, err := p.parseMany(schema.AnyOf)
		if err != nil {
			return nil, errors.Wrapf(err, "anyOf")
		}

		return hook(&Schema{
			AnyOf: schemas,
			// Object validators
			MaxProperties: schema.MaxProperties,
			MinProperties: schema.MinProperties,
			// Array validators
			MinItems:    schema.MinItems,
			MaxItems:    schema.MaxItems,
			UniqueItems: schema.UniqueItems,
			// Number validators
			Minimum:          schema.Minimum,
			Maximum:          schema.Maximum,
			ExclusiveMinimum: schema.ExclusiveMinimum,
			ExclusiveMaximum: schema.ExclusiveMaximum,
			MultipleOf:       schema.MultipleOf,
			// String validators
			MaxLength: schema.MaxLength,
			MinLength: schema.MinLength,
			Pattern:   schema.Pattern,
		}), nil
	case len(schema.AllOf) > 0:
		schemas, err := p.parseMany(schema.AllOf)
		if err != nil {
			return nil, errors.Wrapf(err, "allOf")
		}

		return hook(&Schema{AllOf: schemas}), nil
	}

	// Try to infer schema type from properties.
	if p.inferTypes && schema.Type == "" {
		switch {
		case schema.AdditionalProperties != nil ||
			schema.MaxProperties != nil ||
			schema.MinProperties != nil:
			schema.Type = "object"

		case schema.Items != nil ||
			schema.UniqueItems ||
			schema.MaxItems != nil ||
			schema.MinItems != nil:
			schema.Type = "array"

		case schema.Maximum != nil ||
			schema.Minimum != nil ||
			schema.ExclusiveMinimum ||
			schema.ExclusiveMaximum || // FIXME(tdakkota): check for existence instead of true?
			schema.MultipleOf != nil:
			schema.Type = "number"

		case schema.MaxLength != nil ||
			schema.MinLength != nil ||
			schema.Pattern != "":
			schema.Type = "string"
		}
	}

	switch schema.Type {
	case "object":
		if schema.Items != nil {
			return nil, errors.New("object cannot contain 'items' field")
		}
		required := func(name string) bool {
			for _, p := range schema.Required {
				if p == name {
					return true
				}
			}
			return false
		}

		s := hook(&Schema{
			Type:          Object,
			MaxProperties: schema.MaxProperties,
			MinProperties: schema.MinProperties,
		})

		if ap := schema.AdditionalProperties; ap != nil {
			s.AdditionalProperties = true
			if !ap.Bool {
				item, err := p.Parse(&ap.Schema)
				if err != nil {
					return nil, errors.Wrapf(err, "additionalProperties")
				}
				s.Item = item
			}
		}

		for _, propSpec := range schema.Properties {
			prop, err := p.Parse(propSpec.Schema)
			if err != nil {
				return nil, errors.Wrapf(err, "%s", propSpec.Name)
			}

			s.Properties = append(s.Properties, Property{
				Name:        propSpec.Name,
				Description: propSpec.Schema.Description,
				Schema:      prop,
				Required:    required(propSpec.Name),
			})
		}
		return s, nil

	case "array":
		array := hook(&Schema{
			Type:        Array,
			MinItems:    schema.MinItems,
			MaxItems:    schema.MaxItems,
			UniqueItems: schema.UniqueItems,
		})
		if schema.Items == nil {
			// Keep items nil, we will use ir.Any for it.
			return array, nil
		}
		if len(schema.Properties) > 0 {
			return nil, errors.New("array cannot contain properties")
		}

		item, err := p.Parse(schema.Items)
		if err != nil {
			return nil, errors.Wrap(err, "item")
		}

		array.Item = item
		return array, nil

	case "number", "integer":
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		return hook(&Schema{
			Type:             SchemaType(schema.Type),
			Format:           Format(schema.Format),
			Minimum:          schema.Minimum,
			Maximum:          schema.Maximum,
			ExclusiveMinimum: schema.ExclusiveMinimum,
			ExclusiveMaximum: schema.ExclusiveMaximum,
			MultipleOf:       schema.MultipleOf,
		}), nil

	case "boolean":
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		return hook(&Schema{
			Type:   Boolean,
			Format: Format(schema.Format),
		}), nil

	case "string":
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		return hook(&Schema{
			Type:      String,
			Format:    Format(schema.Format),
			MaxLength: schema.MaxLength,
			MinLength: schema.MinLength,
			Pattern:   schema.Pattern,
		}), nil

	case "":
		return hook(&Schema{}), nil

	default:
		return nil, errors.Errorf("unexpected schema type: %q", schema.Type)
	}
}

func (p *Parser) parseMany(schemas []*RawSchema) ([]*Schema, error) {
	result := make([]*Schema, 0, len(schemas))
	for i, schema := range schemas {
		s, err := p.Parse(schema)
		if err != nil {
			return nil, errors.Wrapf(err, "[%d]", i)
		}

		result = append(result, s)
	}

	return result, nil
}

func (p *Parser) resolve(ref string) (*Schema, error) {
	if s, ok := p.refcache[ref]; ok {
		return s, nil
	}

	raw, err := p.resolver.ResolveReference(ref)
	if err != nil {
		return nil, errors.Wrapf(err, "resolve reference %q", ref)
	}

	return p.parse(raw, func(s *Schema) *Schema {
		s.Ref = ref
		p.refcache[ref] = s
		return p.extendInfo(raw, s)
	})
}

func (p *Parser) extendInfo(schema *RawSchema, s *Schema) *Schema {
	s.Description = schema.Description
	s.AddExample(schema.Example)

	// Workaround: handle nullable enums correctly.
	//
	// Notice that nullable enum requires `null` in value list.
	//
	// See https://github.com/OAI/OpenAPI-Specification/blob/main/proposals/2019-10-31-Clarify-Nullable.md#if-a-schema-specifies-nullable-true-and-enum-1-2-3-does-that-schema-allow-null-values-see-1900.
	if len(s.Enum) < 1 {
		s.Nullable = schema.Nullable
	} else {
		// Check that enum contains `null` value.
		for _, v := range s.Enum {
			if v == nil {
				s.Nullable = true
				break
			}
		}
		// Filter all `null`s.
		if s.Nullable {
			n := 0
			for _, v := range s.Enum {
				if v != nil {
					s.Enum[n] = v
					n++
				}
			}
			s.Enum = s.Enum[:n]
		}
	}
	if d := schema.Discriminator; d != nil {
		s.Discriminator = &Discriminator{
			PropertyName: d.PropertyName,
			Mapping:      d.Mapping,
		}
	}
	return s
}

func validateTypeFormat(typ, format string) error {
	formats := map[string]map[string]struct{}{
		"object":  {},
		"array":   {},
		"boolean": {},
		"integer": {
			"int32": {},
			"int64": {},
		},
		"number": {
			"float":  {},
			"double": {},
			"int32":  {},
			"int64":  {},
		},
		"string": {
			"byte":      {},
			"date-time": {},
			"date":      {},
			"time":      {},
			"duration":  {},
			"uuid":      {},
			"ipv4":      {},
			"ipv6":      {},
			"ip":        {},
			"uri":       {},
			"password":  {},
		},
	}

	if _, ok := formats[typ]; !ok {
		return errors.Errorf("unexpected type: %q", typ)
	}

	if format == "" {
		return nil
	}

	if _, ok := formats[typ][format]; !ok {
		if typ == "string" {
			return nil // Ignore unknown string formats.
		}

		return errors.Errorf("unexpected %s format: %q", typ, format)
	}

	return nil
}
