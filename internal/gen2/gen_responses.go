package gen

import (
	"net/http"

	ast "github.com/ogen-go/ogen/internal/ast2"
	"github.com/ogen-go/ogen/internal/ir"
	"golang.org/x/xerrors"
)

func (g *Generator) generateResponses(name string, responses *ast.MethodResponse) (*ir.Response, error) {
	result := &ir.Response{
		Spec:       responses,
		StatusCode: map[int]*ir.StatusResponse{},
	}

	for code, resp := range responses.StatusCode {
		resp, err := g.response2IR(pascal(name, http.StatusText(code)), resp)
		if err != nil {
			return nil, xerrors.Errorf("%d: %w", code, err)
		}

		result.StatusCode[code] = resp
	}

	if def := responses.Default; def != nil {
		resp, err := g.response2IR(name+"Default", def)
		if err != nil {
			return nil, xerrors.Errorf("default: %w", err)
		}

		for contentType, typ := range resp.Contents {
			resp.Contents[contentType] = g.wrapStatusCode(typ)
		}

		if noc := resp.NoContent; noc != nil {
			resp.NoContent = g.wrapStatusCode(noc)
		}

		result.Default = resp
	}

	return result, nil
}

func (g *Generator) response2IR(name string, resp *ast.Response) (*ir.StatusResponse, error) {
	if len(resp.Contents) == 0 {
		return &ir.StatusResponse{
			NoContent: &ir.Type{
				Kind: ir.KindStruct,
				Name: name,
			},
			Spec: resp,
		}, nil
	}

	types := make(map[string]*ir.Type)
	for contentType, schema := range resp.Contents {
		typ, err := g.generateSchema(pascal(name, contentType), schema)
		if err != nil {
			return nil, xerrors.Errorf("contents: %s: %w", contentType, err)
		}

		types[contentType] = typ
	}

	return &ir.StatusResponse{
		Contents: types,
		Spec:     resp,
	}, nil
}

func (g *Generator) wrapStatusCode(typ *ir.Type) *ir.Type {
	if !typ.Is(ir.KindStruct, ir.KindAlias) {
		panic("unreachable")
	}

	return &ir.Type{
		Kind: ir.KindStruct,
		Name: typ.Name + "Default",
		Fields: []*ir.StructField{
			{
				Name: "StatusCode",
				Type: ir.Primitive(ir.Int, nil),
				Tag:  "-",
			},
			{
				Name: "Response",
				Type: typ,
				Tag:  "-",
			},
		},
	}
}
