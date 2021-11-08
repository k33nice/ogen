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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	serverURL *url.URL
	cfg       config
	requests  metric.Int64Counter
	errors    metric.Int64Counter
	duration  metric.Int64Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.NewInt64Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// CreateSnapshot implements createSnapshot operation.
func (c *Client) CreateSnapshot(ctx context.Context, request SnapshotCreateParams) (res CreateSnapshotRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `CreateSnapshot`,
		trace.WithAttributes(otelogen.OperationID(`createSnapshot`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeCreateSnapshotRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/snapshot/create"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreateSnapshotResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateSyncAction implements createSyncAction operation.
func (c *Client) CreateSyncAction(ctx context.Context, request InstanceActionInfo) (res CreateSyncActionRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `CreateSyncAction`,
		trace.WithAttributes(otelogen.OperationID(`createSyncAction`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeCreateSyncActionRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/actions"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreateSyncActionResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DescribeBalloonConfig implements describeBalloonConfig operation.
func (c *Client) DescribeBalloonConfig(ctx context.Context) (res DescribeBalloonConfigRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `DescribeBalloonConfig`,
		trace.WithAttributes(otelogen.OperationID(`describeBalloonConfig`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/balloon"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDescribeBalloonConfigResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DescribeBalloonStats implements describeBalloonStats operation.
func (c *Client) DescribeBalloonStats(ctx context.Context) (res DescribeBalloonStatsRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `DescribeBalloonStats`,
		trace.WithAttributes(otelogen.OperationID(`describeBalloonStats`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/balloon/statistics"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDescribeBalloonStatsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DescribeInstance implements describeInstance operation.
func (c *Client) DescribeInstance(ctx context.Context) (res DescribeInstanceRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `DescribeInstance`,
		trace.WithAttributes(otelogen.OperationID(`describeInstance`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDescribeInstanceResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetExportVmConfig implements getExportVmConfig operation.
func (c *Client) GetExportVmConfig(ctx context.Context) (res GetExportVmConfigRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `GetExportVmConfig`,
		trace.WithAttributes(otelogen.OperationID(`getExportVmConfig`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/vm/config"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeGetExportVmConfigResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetMachineConfiguration implements getMachineConfiguration operation.
func (c *Client) GetMachineConfiguration(ctx context.Context) (res GetMachineConfigurationRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `GetMachineConfiguration`,
		trace.WithAttributes(otelogen.OperationID(`getMachineConfiguration`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/machine-config"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeGetMachineConfigurationResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LoadSnapshot implements loadSnapshot operation.
func (c *Client) LoadSnapshot(ctx context.Context, request SnapshotLoadParams) (res LoadSnapshotRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `LoadSnapshot`,
		trace.WithAttributes(otelogen.OperationID(`loadSnapshot`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeLoadSnapshotRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/snapshot/load"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeLoadSnapshotResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

func (c *Client) MmdsConfigPut(ctx context.Context, request MmdsConfig) (res MmdsConfigPutRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `MmdsConfigPut`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeMmdsConfigPutRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/mmds/config"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsConfigPutResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

func (c *Client) MmdsGet(ctx context.Context) (res MmdsGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `MmdsGet`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/mmds"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

func (c *Client) MmdsPatch(ctx context.Context, request MmdsPatchReq) (res MmdsPatchRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `MmdsPatch`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeMmdsPatchRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/mmds"

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsPatchResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

func (c *Client) MmdsPut(ctx context.Context, request MmdsPutReq) (res MmdsPutRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `MmdsPut`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeMmdsPutRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/mmds"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsPutResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchBalloon implements patchBalloon operation.
func (c *Client) PatchBalloon(ctx context.Context, request BalloonUpdate) (res PatchBalloonRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PatchBalloon`,
		trace.WithAttributes(otelogen.OperationID(`patchBalloon`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePatchBalloonRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/balloon"

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchBalloonResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchBalloonStatsInterval implements patchBalloonStatsInterval operation.
func (c *Client) PatchBalloonStatsInterval(ctx context.Context, request BalloonStatsUpdate) (res PatchBalloonStatsIntervalRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PatchBalloonStatsInterval`,
		trace.WithAttributes(otelogen.OperationID(`patchBalloonStatsInterval`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePatchBalloonStatsIntervalRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/balloon/statistics"

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchBalloonStatsIntervalResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchGuestDriveByID implements patchGuestDriveByID operation.
func (c *Client) PatchGuestDriveByID(ctx context.Context, request PartialDrive, params PatchGuestDriveByIDParams) (res PatchGuestDriveByIDRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PatchGuestDriveByID`,
		trace.WithAttributes(otelogen.OperationID(`patchGuestDriveByID`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePatchGuestDriveByIDRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/drives/"
	{
		// Encode "drive_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "drive_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.DriveID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchGuestDriveByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchGuestNetworkInterfaceByID implements patchGuestNetworkInterfaceByID operation.
func (c *Client) PatchGuestNetworkInterfaceByID(ctx context.Context, request PartialNetworkInterface, params PatchGuestNetworkInterfaceByIDParams) (res PatchGuestNetworkInterfaceByIDRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PatchGuestNetworkInterfaceByID`,
		trace.WithAttributes(otelogen.OperationID(`patchGuestNetworkInterfaceByID`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePatchGuestNetworkInterfaceByIDRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/network-interfaces/"
	{
		// Encode "iface_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "iface_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.IfaceID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchGuestNetworkInterfaceByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchMachineConfiguration implements patchMachineConfiguration operation.
func (c *Client) PatchMachineConfiguration(ctx context.Context, request MachineConfiguration) (res PatchMachineConfigurationRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PatchMachineConfiguration`,
		trace.WithAttributes(otelogen.OperationID(`patchMachineConfiguration`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePatchMachineConfigurationRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/machine-config"

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchMachineConfigurationResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchVm implements patchVm operation.
func (c *Client) PatchVm(ctx context.Context, request VM) (res PatchVmRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PatchVm`,
		trace.WithAttributes(otelogen.OperationID(`patchVm`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePatchVmRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/vm"

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchVmResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutBalloon implements putBalloon operation.
func (c *Client) PutBalloon(ctx context.Context, request Balloon) (res PutBalloonRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutBalloon`,
		trace.WithAttributes(otelogen.OperationID(`putBalloon`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutBalloonRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/balloon"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutBalloonResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestBootSource implements putGuestBootSource operation.
func (c *Client) PutGuestBootSource(ctx context.Context, request BootSource) (res PutGuestBootSourceRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutGuestBootSource`,
		trace.WithAttributes(otelogen.OperationID(`putGuestBootSource`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutGuestBootSourceRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/boot-source"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestBootSourceResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestDriveByID implements putGuestDriveByID operation.
func (c *Client) PutGuestDriveByID(ctx context.Context, request Drive, params PutGuestDriveByIDParams) (res PutGuestDriveByIDRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutGuestDriveByID`,
		trace.WithAttributes(otelogen.OperationID(`putGuestDriveByID`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutGuestDriveByIDRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/drives/"
	{
		// Encode "drive_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "drive_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.DriveID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestDriveByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestNetworkInterfaceByID implements putGuestNetworkInterfaceByID operation.
func (c *Client) PutGuestNetworkInterfaceByID(ctx context.Context, request NetworkInterface, params PutGuestNetworkInterfaceByIDParams) (res PutGuestNetworkInterfaceByIDRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutGuestNetworkInterfaceByID`,
		trace.WithAttributes(otelogen.OperationID(`putGuestNetworkInterfaceByID`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutGuestNetworkInterfaceByIDRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/network-interfaces/"
	{
		// Encode "iface_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "iface_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.IfaceID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestNetworkInterfaceByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestVsock implements putGuestVsock operation.
func (c *Client) PutGuestVsock(ctx context.Context, request Vsock) (res PutGuestVsockRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutGuestVsock`,
		trace.WithAttributes(otelogen.OperationID(`putGuestVsock`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutGuestVsockRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/vsock"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestVsockResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutLogger implements putLogger operation.
func (c *Client) PutLogger(ctx context.Context, request Logger) (res PutLoggerRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutLogger`,
		trace.WithAttributes(otelogen.OperationID(`putLogger`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutLoggerRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/logger"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutLoggerResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutMachineConfiguration implements putMachineConfiguration operation.
func (c *Client) PutMachineConfiguration(ctx context.Context, request MachineConfiguration) (res PutMachineConfigurationRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutMachineConfiguration`,
		trace.WithAttributes(otelogen.OperationID(`putMachineConfiguration`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutMachineConfigurationRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/machine-config"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutMachineConfigurationResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutMetrics implements putMetrics operation.
func (c *Client) PutMetrics(ctx context.Context, request Metrics) (res PutMetricsRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PutMetrics`,
		trace.WithAttributes(otelogen.OperationID(`putMetrics`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePutMetricsRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/metrics"

	r := ht.NewRequest(ctx, "PUT", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutMetricsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
