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

func decodeDataGetFormatResponse(resp *http.Response, span trace.Span) (res string, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response string
			if err := func() error {
				v, err := d.Str()
				response = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeErrorGetResponse(resp *http.Response, span trace.Span) (res ErrorStatusCode, err error) {
	switch resp.StatusCode {
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeFoobarGetResponse(resp *http.Response, span trace.Span) (res FoobarGetRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &NotFound{}, nil
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeFoobarPostResponse(resp *http.Response, span trace.Span) (res FoobarPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &NotFound{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeFoobarPutResponse(resp *http.Response, span trace.Span) (res FoobarPutDefStatusCode, err error) {
	switch resp.StatusCode {
	default:
		return FoobarPutDefStatusCode{
			StatusCode: resp.StatusCode,
		}, nil
	}
}

func decodeGetHeaderResponse(resp *http.Response, span trace.Span) (res Hash, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Hash
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeOneofBugResponse(resp *http.Response, span trace.Span) (res OneofBugOK, err error) {
	switch resp.StatusCode {
	case 200:
		return OneofBugOK{}, nil
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetCreateResponse(resp *http.Response, span trace.Span) (res Pet, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetFriendsNamesByIDResponse(resp *http.Response, span trace.Span) (res []string, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response []string
			if err := func() error {
				response = nil
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					response = append(response, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetGetResponse(resp *http.Response, span trace.Span) (res PetGetRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetGetDef
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &PetGetDefStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePetGetAvatarByIDResponse(resp *http.Response, span trace.Span) (res PetGetAvatarByIDRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/octet-stream":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			return &PetGetAvatarByIDOK{
				Data: bytes.NewReader(b),
			}, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &NotFound{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePetGetByNameResponse(resp *http.Response, span trace.Span) (res Pet, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetNameByIDResponse(resp *http.Response, span trace.Span) (res string, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response string
			if err := func() error {
				v, err := d.Str()
				response = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetUpdateNameAliasPostResponse(resp *http.Response, span trace.Span) (res PetUpdateNameAliasPostDefStatusCode, err error) {
	switch resp.StatusCode {
	default:
		return PetUpdateNameAliasPostDefStatusCode{
			StatusCode: resp.StatusCode,
		}, nil
	}
}

func decodePetUpdateNamePostResponse(resp *http.Response, span trace.Span) (res PetUpdateNamePostDefStatusCode, err error) {
	switch resp.StatusCode {
	default:
		return PetUpdateNamePostDefStatusCode{
			StatusCode: resp.StatusCode,
		}, nil
	}
}

func decodePetUploadAvatarByIDResponse(resp *http.Response, span trace.Span) (res PetUploadAvatarByIDRes, err error) {
	switch resp.StatusCode {
	case 200:
		return &PetUploadAvatarByIDOK{}, nil
	case 404:
		return &NotFound{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeRecursiveMapGetResponse(resp *http.Response, span trace.Span) (res RecursiveMap, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response RecursiveMap
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeTestObjectQueryParameterResponse(resp *http.Response, span trace.Span) (res TestObjectQueryParameterOK, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response TestObjectQueryParameterOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}
