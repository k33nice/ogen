// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
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
	"time"

	"github.com/go-chi/chi/v5"
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
	_ = chi.Context{}
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
)

// WriteJSON implements json.Marshaler.
func (s DeletePetNoContent) WriteJSON(e *json.Encoder) {
	e.ObjStart()
	more := json.NewMore(e)
	defer more.Reset()
	e.ObjEnd()
}

// ReadJSON reads DeletePetNoContent from json stream.
func (s *DeletePetNoContent) ReadJSON(d *json.Decoder) error {
	if s == nil {
		return fmt.Errorf(`invalid: unable to decode DeletePetNoContent to nil`)
	}
	return d.ObjBytes(func(d *json.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// WriteJSON implements json.Marshaler.
func (s Error) WriteJSON(e *json.Encoder) {
	e.ObjStart()
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	e.ObjField("code")
	e.Int32(s.Code)
	more.More()
	e.ObjField("message")
	e.Str(s.Message)
	e.ObjEnd()
}

// ReadJSON reads Error from json stream.
func (s *Error) ReadJSON(d *json.Decoder) error {
	if s == nil {
		return fmt.Errorf(`invalid: unable to decode Error to nil`)
	}
	return d.ObjBytes(func(d *json.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			v, err := d.Int32()
			s.Code = int32(v)
			if err != nil {
				return err
			}
		case "message":
			v, err := d.Str()
			s.Message = string(v)
			if err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// WriteJSON implements json.Marshaler.
func (s ErrorStatusCode) WriteJSON(e *json.Encoder) {
	e.ObjStart()
	more := json.NewMore(e)
	defer more.Reset()
	e.ObjEnd()
}

// ReadJSON reads ErrorStatusCode from json stream.
func (s *ErrorStatusCode) ReadJSON(d *json.Decoder) error {
	if s == nil {
		return fmt.Errorf(`invalid: unable to decode ErrorStatusCode to nil`)
	}
	return d.ObjBytes(func(d *json.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// WriteJSON implements json.Marshaler.
func (s NewPet) WriteJSON(e *json.Encoder) {
	e.ObjStart()
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	e.ObjField("name")
	e.Str(s.Name)
	if s.Tag.Set {
		more.More()
		e.ObjField("tag")
		s.Tag.WriteJSON(e)
	}
	e.ObjEnd()
}

// ReadJSON reads NewPet from json stream.
func (s *NewPet) ReadJSON(d *json.Decoder) error {
	if s == nil {
		return fmt.Errorf(`invalid: unable to decode NewPet to nil`)
	}
	return d.ObjBytes(func(d *json.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			v, err := d.Str()
			s.Name = string(v)
			if err != nil {
				return err
			}
		case "tag":
			s.Tag.Reset()
			if err := s.Tag.ReadJSON(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// WriteJSON writes json value of string to json stream.
func (o OptString) WriteJSON(e *json.Encoder) {
	e.Str(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptString) ReadJSON(d *json.Decoder) error {
	if o == nil {
		return fmt.Errorf(`invalid: unable to decode OptString to nil`)
	}
	switch d.Next() {
	case json.String:
		o.Set = true
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = string(v)
		return nil
	default:
		return fmt.Errorf("unexpected type %q while reading OptString", d.Next())
	}
}
