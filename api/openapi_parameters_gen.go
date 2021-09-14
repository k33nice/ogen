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

type FoobarGetParams struct {
	Query FoobarGetQueryParams
}

type FoobarGetQueryParams struct {
	InlinedParam int64
	Skip         int32
}

func ParseFoobarGetParams(r *http.Request) (FoobarGetParams, error) {
	var params FoobarGetParams
	{
		param := r.URL.Query().Get("inlinedParam")
		if len(param) == 0 {
			return params, fmt.Errorf("Query param 'inlinedParam' is empty")
		}

		v, err := conv.ToInt64(param)
		if err != nil {
			return params, fmt.Errorf("parse Query param 'inlinedParam': %w", err)
		}

		params.Query.InlinedParam = v
	}
	{
		param := r.URL.Query().Get("skip")
		if len(param) == 0 {
			return params, fmt.Errorf("Query param 'skip' is empty")
		}

		v, err := conv.ToInt32(param)
		if err != nil {
			return params, fmt.Errorf("parse Query param 'skip': %w", err)
		}

		params.Query.Skip = v
	}

	return params, nil
}

type PetGetParams struct {
	Cookie PetGetCookieParams
	Header PetGetHeaderParams
	Query  PetGetQueryParams
}

type PetGetCookieParams struct {
	Token string
}

type PetGetHeaderParams struct {
	XScope []string
}

type PetGetQueryParams struct {
	PetID int64
}

func ParsePetGetParams(r *http.Request) (PetGetParams, error) {
	var params PetGetParams
	{
		c, err := r.Cookie("token")
		if err != nil {
			return params, fmt.Errorf("Cookie param 'token': %w", err)
		}

		param := c.Value
		if len(param) == 0 {
			return params, fmt.Errorf("Cookie param 'token' is empty")
		}

		v, err := conv.ToString(param)
		if err != nil {
			return params, fmt.Errorf("parse Cookie param 'token': %w", err)
		}

		params.Cookie.Token = v
	}
	{
		param := r.Header.Values("x-scope")
		if len(param) == 0 {
			return params, fmt.Errorf("Header param 'x-scope' is empty")
		}

		v, err := conv.ToStringArray(param)
		if err != nil {
			return params, fmt.Errorf("parse Header param 'x-scope': %w", err)
		}

		params.Header.XScope = v
	}
	{
		param := r.URL.Query().Get("petID")
		if len(param) == 0 {
			return params, fmt.Errorf("Query param 'petID' is empty")
		}

		v, err := conv.ToInt64(param)
		if err != nil {
			return params, fmt.Errorf("parse Query param 'petID': %w", err)
		}

		params.Query.PetID = v
	}

	return params, nil
}

type PetNameGetParams struct {
	Path PetNameGetPathParams
}

type PetNameGetPathParams struct {
	Name string
}

func ParsePetNameGetParams(r *http.Request) (PetNameGetParams, error) {
	var params PetNameGetParams
	{
		param := chi.URLParam(r, "name")
		if len(param) == 0 {
			return params, fmt.Errorf("Path param 'name' is empty")
		}

		v, err := conv.ToString(param)
		if err != nil {
			return params, fmt.Errorf("parse Path param 'name': %w", err)
		}

		params.Path.Name = v
	}

	return params, nil
}
