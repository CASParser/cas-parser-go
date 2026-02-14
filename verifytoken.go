// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/cas-parser-go/internal/apijson"
	"github.com/stainless-sdks/cas-parser-go/internal/requestconfig"
	"github.com/stainless-sdks/cas-parser-go/option"
	"github.com/stainless-sdks/cas-parser-go/packages/respjson"
)

// VerifyTokenService contains methods and other services that help with
// interacting with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerifyTokenService] method instead.
type VerifyTokenService struct {
	Options []option.RequestOption
}

// NewVerifyTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerifyTokenService(opts ...option.RequestOption) (r VerifyTokenService) {
	r = VerifyTokenService{}
	r.Options = opts
	return
}

// Verify an access token and check if it's still valid. Useful for debugging token
// issues.
func (r *VerifyTokenService) Verify(ctx context.Context, opts ...option.RequestOption) (res *VerifyTokenVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/verify-token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type VerifyTokenVerifyResponse struct {
	// Error message (only shown if invalid)
	Error string `json:"error"`
	// Masked API key (only shown if valid)
	MaskedAPIKey string `json:"masked_api_key"`
	// Whether the token is valid
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error        respjson.Field
		MaskedAPIKey respjson.Field
		Valid        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyTokenVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *VerifyTokenVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
