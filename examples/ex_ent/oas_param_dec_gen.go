// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func decodeCreatePetCategoriesParams(args map[string]string, r *http.Request) (CreatePetCategoriesParams, error) {
	var (
		params CreatePetCategoriesParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}

func decodeCreatePetFriendsParams(args map[string]string, r *http.Request) (CreatePetFriendsParams, error) {
	var (
		params CreatePetFriendsParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}

func decodeCreatePetOwnerParams(args map[string]string, r *http.Request) (CreatePetOwnerParams, error) {
	var (
		params CreatePetOwnerParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}

func decodeDeletePetParams(args map[string]string, r *http.Request) (DeletePetParams, error) {
	var (
		params DeletePetParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}

func decodeDeletePetOwnerParams(args map[string]string, r *http.Request) (DeletePetOwnerParams, error) {
	var (
		params DeletePetOwnerParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}

func decodeListPetParams(args map[string]string, r *http.Request) (ListPetParams, error) {
	var (
		params    ListPetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsPageVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: page: parse`)
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsItemsPerPageVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: itemsPerPage: parse`)
			}
		}
	}
	return params, nil
}

func decodeListPetCategoriesParams(args map[string]string, r *http.Request) (ListPetCategoriesParams, error) {
	var (
		params    ListPetCategoriesParams
		queryArgs = r.URL.Query()
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsPageVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: page: parse`)
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsItemsPerPageVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: itemsPerPage: parse`)
			}
		}
	}
	return params, nil
}

func decodeListPetFriendsParams(args map[string]string, r *http.Request) (ListPetFriendsParams, error) {
	var (
		params    ListPetFriendsParams
		queryArgs = r.URL.Query()
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsPageVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: page: parse`)
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsItemsPerPageVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: itemsPerPage: parse`)
			}
		}
	}
	return params, nil
}

func decodeReadPetParams(args map[string]string, r *http.Request) (ReadPetParams, error) {
	var (
		params ReadPetParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}

func decodeReadPetOwnerParams(args map[string]string, r *http.Request) (ReadPetOwnerParams, error) {
	var (
		params ReadPetOwnerParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}

func decodeUpdatePetParams(args map[string]string, r *http.Request) (UpdatePetParams, error) {
	var (
		params UpdatePetParams
	)
	// Decode path: id.
	{
		param := args["id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: id: not specified`)
		}
	}
	return params, nil
}
