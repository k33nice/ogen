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
	jsoniter "github.com/json-iterator/go"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
	"github.com/ogen-go/ogen/types"
	"github.com/ogen-go/ogen/uri"
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
	_ = types.Date{}
	_ = jsoniter.Config{}
)

func (s HelloWorld) WriteJSON(js *jsoniter.Stream) {
	js.WriteObjectStart()
	js.WriteObjectField("message")
	js.WriteString(s.Message)
	js.WriteObjectEnd()
}

func (s HelloWorld) WriteFieldJSON(k string, js *jsoniter.Stream) {
	js.WriteObjectField(k)
	s.WriteJSON(js)
}

func (s WorldObject) WriteJSON(js *jsoniter.Stream) {
	js.WriteObjectStart()
	js.WriteObjectField("id")
	js.WriteInt64(s.ID)
	js.WriteObjectField("randomNumber")
	js.WriteInt64(s.RandomNumber)
	js.WriteObjectEnd()
}

func (s WorldObject) WriteFieldJSON(k string, js *jsoniter.Stream) {
	js.WriteObjectField(k)
	s.WriteJSON(js)
}
