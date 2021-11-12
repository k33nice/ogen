// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
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

func decodeFoobarPostRequest(r *http.Request, span trace.Span) (req Pet, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		buf := getBuf()
		defer putBuf(buf)
		if _, err := io.Copy(buf, r.Body); err != nil {
			return req, err
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetCreateRequest(r *http.Request, span trace.Span) (req Pet, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		buf := getBuf()
		defer putBuf(buf)
		if _, err := io.Copy(buf, r.Body); err != nil {
			return req, err
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUpdateNameAliasPostRequest(r *http.Request, span trace.Span) (req PetName, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PetName
		buf := getBuf()
		defer putBuf(buf)
		if _, err := io.Copy(buf, r.Body); err != nil {
			return req, err
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			{
				var unwrapped string
				v, err := d.Str()
				unwrapped = string(v)
				if err != nil {
					return err
				}
				request = PetName(unwrapped)
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUpdateNamePostRequest(r *http.Request, span trace.Span) (req string, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request string
		buf := getBuf()
		defer putBuf(buf)
		if _, err := io.Copy(buf, r.Body); err != nil {
			return req, err
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			v, err := d.Str()
			request = string(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := (validate.String{
				MinLength:    6,
				MinLengthSet: true,
				MaxLength:    0,
				MaxLengthSet: false,
				Email:        false,
				Hostname:     false,
				Regex:        nil,
			}).Validate(string(request)); err != nil {
				return errors.Wrap(err, "string")
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUploadAvatarByIDRequest(r *http.Request, span trace.Span) (req Stream, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/octet-stream":
		return Stream{Data: r.Body}, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}
