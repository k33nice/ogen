package gen

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/internal/ir"
	"github.com/ogen-go/ogen/internal/oas"
)

func (g *Generator) generateRequest(opName string, body *oas.RequestBody) (*ir.Request, error) {
	name := opName + "Req"

	contents, err := g.generateContents(name, body.Contents)
	if err != nil {
		return nil, errors.Wrap(err, "contents")
	}

	if !body.Required {
		// NOTE:
		// In case where requestBody has multiple content types
		// we can try to wrap sum-type interface with Optional[T]
		// instead of wrapping each content type.
		for contentType, typ := range contents {
			// TODO: Support optional streams?
			if typ.Is(ir.KindStream) {
				continue
			}

			optionalTyp := ir.Generic(genericPostfix(typ.Go()), typ, ir.GenericVariant{
				Optional: true,
			})

			g.saveType(optionalTyp)
			contents[contentType] = optionalTyp
		}
	}

	if len(contents) == 1 {
		for _, t := range contents {
			return &ir.Request{
				Type:     t,
				Contents: contents,
				Spec:     body,
			}, nil
		}
	}

	iface := ir.Interface(name)
	iface.AddMethod(camel(name))
	g.saveIface(iface)
	for contentType, t := range contents {
		switch t.Kind {
		case ir.KindPrimitive, ir.KindArray:
			// Primitive types cannot have methods, wrap it with alias.
			t = ir.Alias(pascal(name, string(contentType)), t)
			contents[contentType] = t
			g.saveType(t)
		default:
		}

		t.Implement(iface)
	}

	return &ir.Request{
		Type:     iface,
		Contents: contents,
		Spec:     body,
	}, nil
}
