// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
	_ = math.Mod
	_ = validate.Int{}
)

// WriteJSON implements json.Marshaler.
func (s HelloWorld) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes HelloWorld json value to io.Writer.
func (s HelloWorld) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads HelloWorld json value from io.Reader.
func (s *HelloWorld) ReadJSONFrom(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	i := json.NewIterator(json.ConfigDefault)
	i.ResetBytes(data)
	return s.ReadJSON(i)
}

// ReadJSON reads HelloWorld from json stream.
func (s *HelloWorld) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "message":
			s.Message = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s WorldObject) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("id")
	j.WriteInt64(s.ID)
	field.Write("randomNumber")
	j.WriteInt64(s.RandomNumber)
	j.WriteObjectEnd()
}

// WriteJSONTo writes WorldObject json value to io.Writer.
func (s WorldObject) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads WorldObject json value from io.Reader.
func (s *WorldObject) ReadJSONFrom(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	i := json.NewIterator(json.ConfigDefault)
	i.ResetBytes(data)
	return s.ReadJSON(i)
}

// ReadJSON reads WorldObject from json stream.
func (s *WorldObject) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "id":
			s.ID = i.ReadInt64()
			return i.Error == nil
		case "randomNumber":
			s.RandomNumber = i.ReadInt64()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}
