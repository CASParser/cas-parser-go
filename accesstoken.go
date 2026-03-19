// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"net/http"
	"slices"

	"github.com/CASParser/cas-parser-go/internal/apijson"
	"github.com/CASParser/cas-parser-go/internal/requestconfig"
	"github.com/CASParser/cas-parser-go/option"
	"github.com/CASParser/cas-parser-go/packages/param"
	"github.com/CASParser/cas-parser-go/packages/respjson"
)

// Endpoints for managing access tokens for the Portfolio Connect SDK. Use these to
// generate short-lived `at_` prefixed tokens that can be safely passed to frontend
// applications. Access tokens can be used in place of API keys on all v4
// endpoints.
//
// AccessTokenService contains methods and other services that help with
// interacting with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessTokenService] method instead.
type AccessTokenService struct {
	Options []option.RequestOption
}

// NewAccessTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessTokenService(opts ...option.RequestOption) (r AccessTokenService) {
	r = AccessTokenService{}
	r.Options = opts
	return
}

// Generate a short-lived access token from your API key.
//
// **Use this endpoint from your backend** to create tokens that can be safely
// passed to frontend/SDK.
//
// **Legacy path:** `/v1/access-token` (still supported)
//
// Access tokens:
//
// - Are prefixed with `at_` for easy identification
// - Valid for up to 60 minutes
// - Can be used in place of API keys on all v4 endpoints
// - Cannot be used to generate other access tokens
func (r *AccessTokenService) New(ctx context.Context, body AccessTokenNewParams, opts ...option.RequestOption) (res *AccessTokenNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AccessTokenNewResponse struct {
	// The at\_ prefixed access token
	AccessToken string `json:"access_token"`
	// Token validity in seconds
	ExpiresIn int64 `json:"expires_in"`
	// Always "api_key" - token is a drop-in replacement for x-api-key header
	TokenType string `json:"token_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		ExpiresIn   respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessTokenNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessTokenNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTokenNewParams struct {
	// Token validity in minutes (max 60)
	ExpiryMinutes param.Opt[int64] `json:"expiry_minutes,omitzero"`
	paramObj
}

func (r AccessTokenNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessTokenNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessTokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
