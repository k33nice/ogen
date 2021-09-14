// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
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
)

func NewFoobarGetHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		params, err := ParseFoobarGetParams(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.FoobarGet(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func NewFoobarPostHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		request, err := DecodeFoobarPostRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.FoobarPost(r.Context(), request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func NewPetGetHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		params, err := ParsePetGetParams(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.PetGet(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func NewPetPostHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		request, err := DecodePetPostRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.PetPost(r.Context(), request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func NewPetNameGetHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		params, err := ParsePetNameGetParams(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.PetNameGet(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
