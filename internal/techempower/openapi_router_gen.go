// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
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
)

func Register(r chi.Router, s Server) {
	r.MethodFunc("GET", "/cached-worlds", NewCachingHandler(s))
	r.MethodFunc("GET", "/db", NewDBHandler(s))
	r.MethodFunc("GET", "/json", NewJSONHandler(s))
	r.MethodFunc("GET", "/queries", NewQueriesHandler(s))
	r.MethodFunc("GET", "/updates", NewUpdatesHandler(s))
}
