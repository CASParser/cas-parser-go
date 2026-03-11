// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/CASParser/cas-parser-go/internal/apijson"
	"github.com/CASParser/cas-parser-go/internal/requestconfig"
	"github.com/CASParser/cas-parser-go/option"
	"github.com/CASParser/cas-parser-go/packages/param"
	"github.com/CASParser/cas-parser-go/packages/respjson"
)

// Endpoints for fetching CAS documents with instant download. Currently supports
// CDSL via OTP authentication.
//
// CdslFetchService contains methods and other services that help with interacting
// with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCdslFetchService] method instead.
type CdslFetchService struct {
	Options []option.RequestOption
}

// NewCdslFetchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCdslFetchService(opts ...option.RequestOption) (r CdslFetchService) {
	r = CdslFetchService{}
	r.Options = opts
	return
}

// **Step 1 of 2**: Request OTP for CDSL CAS fetch.
//
// This endpoint:
//
// 1. Solves reCAPTCHA automatically (~15-20 seconds)
// 2. Submits login credentials to CDSL portal
// 3. Triggers OTP to user's registered mobile number
//
// After user receives OTP, call `/v4/cdsl/fetch/{session_id}/verify` to complete.
func (r *CdslFetchService) RequestOtp(ctx context.Context, body CdslFetchRequestOtpParams, opts ...option.RequestOption) (res *CdslFetchRequestOtpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/cdsl/fetch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Step 2 of 2**: Verify OTP and retrieve CDSL CAS files.
//
// After successful verification, CAS PDFs are fetched from CDSL portal, uploaded
// to cloud storage, and returned as direct download URLs.
func (r *CdslFetchService) VerifyOtp(ctx context.Context, sessionID string, body CdslFetchVerifyOtpParams, opts ...option.RequestOption) (res *CdslFetchVerifyOtpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v4/cdsl/fetch/%s/verify", url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CdslFetchRequestOtpResponse struct {
	Msg string `json:"msg"`
	// Session ID for verify step
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		SessionID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdslFetchRequestOtpResponse) RawJSON() string { return r.JSON.raw }
func (r *CdslFetchRequestOtpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdslFetchVerifyOtpResponse struct {
	Files  []CdslFetchVerifyOtpResponseFile `json:"files"`
	Msg    string                           `json:"msg"`
	Status string                           `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdslFetchVerifyOtpResponse) RawJSON() string { return r.JSON.raw }
func (r *CdslFetchVerifyOtpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdslFetchVerifyOtpResponseFile struct {
	Filename string `json:"filename"`
	// Direct download URL (cloud storage)
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdslFetchVerifyOtpResponseFile) RawJSON() string { return r.JSON.raw }
func (r *CdslFetchVerifyOtpResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdslFetchRequestOtpParams struct {
	// CDSL BO ID (16 digits)
	BoID string `json:"bo_id" api:"required"`
	// Date of birth (YYYY-MM-DD)
	Dob string `json:"dob" api:"required"`
	// PAN number
	Pan string `json:"pan" api:"required"`
	paramObj
}

func (r CdslFetchRequestOtpParams) MarshalJSON() (data []byte, err error) {
	type shadow CdslFetchRequestOtpParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CdslFetchRequestOtpParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdslFetchVerifyOtpParams struct {
	// OTP received on mobile
	Otp string `json:"otp" api:"required"`
	// Number of monthly statements to fetch (default 6)
	NumPeriods param.Opt[int64] `json:"num_periods,omitzero"`
	paramObj
}

func (r CdslFetchVerifyOtpParams) MarshalJSON() (data []byte, err error) {
	type shadow CdslFetchVerifyOtpParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CdslFetchVerifyOtpParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
