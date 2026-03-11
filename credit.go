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
	"github.com/CASParser/cas-parser-go/packages/respjson"
)

// Endpoints for checking API quota and credits usage. These endpoints help you
// monitor your API usage and remaining quota.
//
// CreditService contains methods and other services that help with interacting
// with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditService] method instead.
type CreditService struct {
	Options []option.RequestOption
}

// NewCreditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCreditService(opts ...option.RequestOption) (r CreditService) {
	r = CreditService{}
	r.Options = opts
	return
}

// Check your remaining API credits and usage for the current billing period.
//
// Returns:
//
// - Number of API calls used and remaining credits
// - Credit limit and reset date
// - List of enabled features for your plan
//
// Credits reset at the start of each billing period.
func (r *CreditService) Check(ctx context.Context, opts ...option.RequestOption) (res *CreditCheckResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/credits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type CreditCheckResponse struct {
	// List of API features enabled for your plan
	EnabledFeatures []string `json:"enabled_features"`
	// Whether the account has unlimited credits
	IsUnlimited bool `json:"is_unlimited"`
	// Total credit limit for billing period
	Limit int64 `json:"limit"`
	// Remaining credits (null if unlimited)
	Remaining float64 `json:"remaining" api:"nullable"`
	// When credits reset (ISO 8601)
	ResetsAt time.Time `json:"resets_at" api:"nullable" format:"date-time"`
	// Number of credits used this billing period
	Used float64 `json:"used"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnabledFeatures respjson.Field
		IsUnlimited     respjson.Field
		Limit           respjson.Field
		Remaining       respjson.Field
		ResetsAt        respjson.Field
		Used            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
