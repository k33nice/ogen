package gen

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/ast"
)

type methodResponse struct {
	Responses map[int]*ast.Response
	Default   *ast.Response
}

func (g *Generator) generateResponses(methodName string, responses ogen.Responses) (*methodResponse, error) {
	result := &methodResponse{
		Responses: map[int]*ast.Response{},
	}

	if len(responses) == 0 {
		return nil, fmt.Errorf("no responses")
	}

	// Iterate over method responses...
	for status, response := range responses {
		// Default response.
		if status == "default" {
			resp, err := g.createDefaultResponse(methodName, response)
			if err != nil {
				return nil, xerrors.Errorf("default: %w", err)
			}

			result.Default = resp
			continue
		}

		statusCode, err := strconv.Atoi(status)
		if err != nil {
			return nil, xerrors.Errorf("invalid status code: '%s'", status)
		}

		if err := func() error {
			// Referenced response.
			if ref := response.Ref; ref != "" {
				// Validate reference & get response component name.
				name, err := componentName(ref)
				if err != nil {
					return err
				}

				// Lookup for response component.
				r, found := g.responses[name]
				if !found {
					return xerrors.Errorf("response by reference '%s' not found", ref)
				}

				result.Responses[statusCode] = r
				return nil
			}

			responseName := pascal(methodName)
			if len(responses) > 1 {
				// Avoid collision with <methodName>Response interface.
				responseName = pascal(methodName, http.StatusText(statusCode))
			}

			resp, err := g.generateResponse(responseName, response)
			if err != nil {
				return err
			}

			result.Responses[statusCode] = resp
			return nil
		}(); err != nil {
			return nil, xerrors.Errorf("%s: %w", status, err)
		}
	}

	return result, nil
}

// createDefaultResponse creates new default response.
func (g *Generator) createDefaultResponse(methodName string, r ogen.Response) (*ast.Response, error) {
	if ref := r.Ref; ref != "" {
		// Validate reference & get response component name.
		name, err := componentName(ref)
		if err != nil {
			return nil, fmt.Errorf("name: %w", err)
		}

		// Lookup for alias response.
		if alias, ok := g.responses[name+"Default"]; ok {
			return alias, nil
		}

		// Lookup for reference response.
		response, found := g.responses[name]
		if !found {
			return nil, xerrors.Errorf("response by reference '%s', not found", ref)
		}

		alias := ast.CreateResponse()
		for contentType, schema := range response.Contents {
			alias.Contents[contentType] = g.wrapStatusCode(schema)
		}
		if schema := response.NoContent; schema != nil {
			response.NoContent = g.wrapStatusCode(schema)
		}

		g.responses[name+"Default"] = alias
		return alias, nil
	}

	// Default response with no contents.
	if len(r.Content) == 0 {
		statusCode := ast.CreateSchemaStruct(methodName + "Default")
		statusCode.Fields = append(statusCode.Fields, ast.SchemaField{
			Name: "StatusCode",
			Type: "int",
			Tag:  "-",
		})
		g.schemas[methodName+"Default"] = statusCode
		return &ast.Response{NoContent: statusCode}, nil
	}

	// Inlined response.
	// Use method name + Default as prefix for response schemas.
	response, err := g.generateResponse(methodName+"Default", r)
	if err != nil {
		return nil, err
	}

	// We need to inject StatusCode field to response structs somehow...
	// Iterate over all responses and create new response schema wrapper:
	//
	// type <WrapperName> struct {
	//     StatusCode int            `json:"-"`
	//     Response   <ResponseType> `json:"-"`
	// }
	for contentType, schema := range response.Contents {
		defaultSchema := g.wrapStatusCode(schema)
		response.Contents[contentType] = defaultSchema
	}
	if schema := response.NoContent; schema != nil {
		response.NoContent = g.wrapStatusCode(schema)
	}

	return response, nil
}

// generateResponse creates new response based on schema definition.
func (g *Generator) generateResponse(rname string, resp ogen.Response) (*ast.Response, error) {
	response := ast.CreateResponse()

	// Response without content.
	// Create empty struct.
	if len(resp.Content) == 0 {
		s := ast.CreateSchemaAlias(rname, "struct{}")
		g.schemas[s.Name] = s
		response.NoContent = s
		return response, nil
	}

	for contentType, media := range resp.Content {
		// Create unique response name.
		name := rname
		if len(resp.Content) > 1 {
			name = pascal(rname, contentType)
		}

		var schema *ast.Schema
		if ref := media.Schema.Ref; ref != "" {
			s, err := g.resolveSchema(ref)
			if err != nil {
				return nil, xerrors.Errorf("content: %s: schema referenced by '%s' not found", contentType, ref)
			}

			schema = s
		} else {
			// Inlined response schema.
			s, err := g.generateSchema(name, media.Schema)
			if xerrors.Is(err, errSkipSchema) {
				continue
			}
			if err != nil {
				return nil, xerrors.Errorf("content: %s: schema: %w", contentType, err)
			}

			schema = s
		}

		if schema.Is(ast.KindPrimitive, ast.KindArray) {
			schema = ast.CreateSchemaAlias(name, schema.Type())
		}

		g.schemas[schema.Name] = schema
		response.Contents[contentType] = schema
	}

	return response, nil
}

// wrapStatusCode wraps provided schema with newtype containing StatusCode field.
//
// Example 1:
//   Schema:
//   type FoobarGetDefaultResponse {
//       Message string `json:"message"`
//       Code    int64  `json:"code"`
//   }
//
//   Wrapper:
//   type FoobarGetDefaultResponseStatusCode {
//       StatusCode int                      `json:"-"`
//       Response   FoobarGetDefaultResponse `json:"-"`
//   }
//
// Example 2:
//   Schema:
//   type FoobarGetDefaultResponse string
//
//   Wrapper:
//   type FoobarGetDefaultResponseStatusCode {
//       StatusCode int    `json:"-"`
//       Response   string `json:"-"`
//   }
//
// TODO: Remove unused schema (Example 2).
func (g *Generator) wrapStatusCode(schema *ast.Schema) *ast.Schema {
	// Use 'StatusCode' postfix for wrapper struct name
	// to avoid name collision with original response schema.
	newSchema := ast.CreateSchemaStruct(schema.Name + "StatusCode")
	newSchema.Fields = []ast.SchemaField{
		{
			Name: "StatusCode",
			Tag:  "-",
			Type: "int",
		},
		{
			Name: "Response",
			Tag:  "-",
			Type: schema.Type(),
		},
	}
	g.schemas[newSchema.Name] = newSchema
	return newSchema
}
