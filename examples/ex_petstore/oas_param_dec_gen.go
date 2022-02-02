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

func decodeListPetsParams(args [0]string, r *http.Request) (ListPetsParams, error) {
	var (
		params    ListPetsParams
		queryArgs = r.URL.Query()
	)
	// Decode query: limit.
	{
		values, ok := queryArgs["limit"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsLimitVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsLimitVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: limit: parse")
			}
		}
	}
	return params, nil
}

func decodeShowPetByIdParams(args [1]string, r *http.Request) (ShowPetByIdParams, error) {
	var (
		params ShowPetByIdParams
	)
	// Decode path: petId.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "petId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.PetId = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: petId: not specified")
		}
	}
	return params, nil
}
