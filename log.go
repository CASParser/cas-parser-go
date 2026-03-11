// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/CASParser/cas-parser-go/internal/apijson"
	"github.com/CASParser/cas-parser-go/internal/requestconfig"
	"github.com/CASParser/cas-parser-go/option"
	"github.com/CASParser/cas-parser-go/packages/param"
	"github.com/CASParser/cas-parser-go/packages/respjson"
)

// Endpoints for checking API quota and credits usage. These endpoints help you
// monitor your API usage and remaining quota.
//
// LogService contains methods and other services that help with interacting with
// the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r LogService) {
	r = LogService{}
	r.Options = opts
	return
}

// Retrieve detailed API usage logs for your account.
//
// Returns a list of API calls with timestamps, features used, status codes, and
// credits consumed. Useful for monitoring usage patterns and debugging.
//
// **Legacy path:** `/logs` (still supported)
func (r *LogService) New(ctx context.Context, body LogNewParams, opts ...option.RequestOption) (res *LogNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get aggregated usage statistics grouped by feature.
//
// Useful for understanding which API features are being used most and tracking
// usage trends.
//
// **Legacy path:** `/logs/summary` (still supported)
func (r *LogService) GetSummary(ctx context.Context, body LogGetSummaryParams, opts ...option.RequestOption) (res *LogGetSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/usage/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type LogNewResponse struct {
	// Number of logs returned
	Count  int64               `json:"count"`
	Logs   []LogNewResponseLog `json:"logs"`
	Status string              `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Logs        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LogNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogNewResponseLog struct {
	// Credits consumed for this request
	Credits float64 `json:"credits"`
	// API feature used
	Feature string `json:"feature"`
	// API endpoint path
	Path string `json:"path"`
	// Unique request identifier
	RequestID string `json:"request_id"`
	// HTTP response status code
	StatusCode int64 `json:"status_code"`
	// When the request was made
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credits     respjson.Field
		Feature     respjson.Field
		Path        respjson.Field
		RequestID   respjson.Field
		StatusCode  respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogNewResponseLog) RawJSON() string { return r.JSON.raw }
func (r *LogNewResponseLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogGetSummaryResponse struct {
	Status  string                       `json:"status"`
	Summary LogGetSummaryResponseSummary `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *LogGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogGetSummaryResponseSummary struct {
	// Usage breakdown by feature
	ByFeature []LogGetSummaryResponseSummaryByFeature `json:"by_feature"`
	// Total credits consumed in the period
	TotalCredits float64 `json:"total_credits"`
	// Total API requests made in the period
	TotalRequests int64 `json:"total_requests"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByFeature     respjson.Field
		TotalCredits  respjson.Field
		TotalRequests respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogGetSummaryResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *LogGetSummaryResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogGetSummaryResponseSummaryByFeature struct {
	// Credits consumed by this feature
	Credits float64 `json:"credits"`
	// API feature name
	Feature string `json:"feature"`
	// Number of requests for this feature
	Requests int64 `json:"requests"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credits     respjson.Field
		Feature     respjson.Field
		Requests    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogGetSummaryResponseSummaryByFeature) RawJSON() string { return r.JSON.raw }
func (r *LogGetSummaryResponseSummaryByFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogNewParams struct {
	// End time filter (ISO 8601). Defaults to now.
	EndTime param.Opt[time.Time] `json:"end_time,omitzero" format:"date-time"`
	// Maximum number of logs to return
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Start time filter (ISO 8601). Defaults to 30 days ago.
	StartTime param.Opt[time.Time] `json:"start_time,omitzero" format:"date-time"`
	paramObj
}

func (r LogNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogGetSummaryParams struct {
	// End time filter (ISO 8601). Defaults to now.
	EndTime param.Opt[time.Time] `json:"end_time,omitzero" format:"date-time"`
	// Start time filter (ISO 8601). Defaults to start of current month.
	StartTime param.Opt[time.Time] `json:"start_time,omitzero" format:"date-time"`
	paramObj
}

func (r LogGetSummaryParams) MarshalJSON() (data []byte, err error) {
	type shadow LogGetSummaryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogGetSummaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
