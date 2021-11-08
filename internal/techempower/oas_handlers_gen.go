// Code generated by ogen, DO NOT EDIT.

package techempower

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

func (s *HTTPServer) HandleCachingRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `Caching`,
		trace.WithAttributes(otelogen.OperationID(`Caching`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCachingParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.Caching(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCachingResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleDBRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DB`,
		trace.WithAttributes(otelogen.OperationID(`DB`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.DB(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDBResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleJSONRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `JSON`,
		trace.WithAttributes(otelogen.OperationID(`json`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.JSON(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeJSONResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleQueriesRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `Queries`,
		trace.WithAttributes(otelogen.OperationID(`Queries`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeQueriesParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.Queries(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeQueriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleUpdatesRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `Updates`,
		trace.WithAttributes(otelogen.OperationID(`Updates`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdatesParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.Updates(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdatesResponse(response, w, span); err != nil {
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
