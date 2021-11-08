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
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-faster/errors"
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

func (s *HTTPServer) HandleFoobarGetRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `FoobarGet`,
		trace.WithAttributes(otelogen.OperationID(`foobarGet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeFoobarGetParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.FoobarGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeFoobarGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleFoobarPostRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `FoobarPost`,
		trace.WithAttributes(otelogen.OperationID(`foobarPost`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeFoobarPostRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.FoobarPost(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeFoobarPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleFoobarPutRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `FoobarPut`,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.FoobarPut(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeFoobarPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetCreateRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetCreate`,
		trace.WithAttributes(otelogen.OperationID(`petCreate`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePetCreateRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetCreate(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetCreateResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetFriendsNamesByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetFriendsNamesByID`,
		trace.WithAttributes(otelogen.OperationID(`petFriendsNamesByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePetFriendsNamesByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetFriendsNamesByID(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetFriendsNamesByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetGetRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetGet`,
		trace.WithAttributes(otelogen.OperationID(`petGet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePetGetParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetGetAvatarByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetGetAvatarByID`,
		trace.WithAttributes(otelogen.OperationID(`petGetAvatarByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePetGetAvatarByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetGetAvatarByID(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetGetAvatarByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetGetByNameRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetGetByName`,
		trace.WithAttributes(otelogen.OperationID(`petGetByName`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePetGetByNameParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetGetByName(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetGetByNameResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetNameByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetNameByID`,
		trace.WithAttributes(otelogen.OperationID(`petNameByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePetNameByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetNameByID(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetNameByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetUpdateNameAliasPostRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetUpdateNameAliasPost`,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePetUpdateNameAliasPostRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetUpdateNameAliasPost(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetUpdateNameAliasPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetUpdateNamePostRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetUpdateNamePost`,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePetUpdateNamePostRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetUpdateNamePost(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetUpdateNamePostResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePetUploadAvatarByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PetUploadAvatarByID`,
		trace.WithAttributes(otelogen.OperationID(`petUploadAvatarByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePetUploadAvatarByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePetUploadAvatarByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PetUploadAvatarByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePetUploadAvatarByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
