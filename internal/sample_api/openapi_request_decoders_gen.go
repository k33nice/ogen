// Code generated by ogen, DO NOT EDIT.

package api

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

func decodeFoobarPostRequest(r *http.Request) (_ *Pet, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			rerr = err
			return
		}

		return &request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePetCreateRequest(r *http.Request) (_ PetCreateRequest, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			rerr = err
			return
		}

		return &request, nil
	case "text/plain":
		var request PetCreateTextPlainRequest
		_ = request
		rerr = fmt.Errorf("text/plain decoder not implemented")
		return
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}
