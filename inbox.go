// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CASParser/cas-parser-go/internal/apijson"
	"github.com/CASParser/cas-parser-go/internal/requestconfig"
	"github.com/CASParser/cas-parser-go/option"
	"github.com/CASParser/cas-parser-go/packages/param"
	"github.com/CASParser/cas-parser-go/packages/respjson"
)

// InboxService contains methods and other services that help with interacting with
// the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxService] method instead.
type InboxService struct {
	Options []option.RequestOption
}

// NewInboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInboxService(opts ...option.RequestOption) (r InboxService) {
	r = InboxService{}
	r.Options = opts
	return
}

// Verify if an `inbox_token` is still valid and check connection status.
//
// Use this to check if the user needs to re-authenticate (e.g., if they revoked
// access in their email provider settings).
func (r *InboxService) CheckConnectionStatus(ctx context.Context, body InboxCheckConnectionStatusParams, opts ...option.RequestOption) (res *InboxCheckConnectionStatusResponse, err error) {
	if !param.IsOmitted(body.XInboxToken) {
		opts = append(opts, option.WithHeader("x-inbox-token", fmt.Sprintf("%v", body.XInboxToken)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v4/inbox/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Initiate OAuth flow to connect user's email inbox.
//
// Returns an `oauth_url` that you should redirect the user to. After
// authorization, they are redirected back to your `redirect_uri` with the
// following query parameters:
//
// **On success:**
//
// - `inbox_token` - Encrypted token to store client-side
// - `email` - Email address of the connected account
// - `state` - Your original state parameter (for CSRF verification)
//
// **On error:**
//
// - `error` - Error code (e.g., `access_denied`, `token_exchange_failed`)
// - `state` - Your original state parameter
//
// **Store the `inbox_token` client-side** and use it for all subsequent inbox API
// calls.
func (r *InboxService) ConnectEmail(ctx context.Context, body InboxConnectEmailParams, opts ...option.RequestOption) (res *InboxConnectEmailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/inbox/connect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Revoke email access and invalidate the token.
//
// This calls the provider's token revocation API (e.g., Google's revoke endpoint)
// to ensure the user's consent is properly removed.
//
// After calling this, the `inbox_token` becomes unusable.
func (r *InboxService) DisconnectEmail(ctx context.Context, body InboxDisconnectEmailParams, opts ...option.RequestOption) (res *InboxDisconnectEmailResponse, err error) {
	if !param.IsOmitted(body.XInboxToken) {
		opts = append(opts, option.WithHeader("x-inbox-token", fmt.Sprintf("%v", body.XInboxToken)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v4/inbox/disconnect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Search the user's email inbox for CAS files from known senders (CAMS, KFintech,
// CDSL, NSDL).
//
// Files are uploaded to temporary cloud storage. **URLs expire in 24 hours.**
//
// Optionally filter by CAS provider and date range.
//
// **Billing:** 0.2 credits per request (charged regardless of success or number of
// files found).
func (r *InboxService) ListCasFiles(ctx context.Context, params InboxListCasFilesParams, opts ...option.RequestOption) (res *InboxListCasFilesResponse, err error) {
	if !param.IsOmitted(params.XInboxToken) {
		opts = append(opts, option.WithHeader("x-inbox-token", fmt.Sprintf("%v", params.XInboxToken)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v4/inbox/cas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type InboxCheckConnectionStatusResponse struct {
	// Whether the token is valid and usable
	Connected bool `json:"connected"`
	// Email address of the connected account
	Email    string `json:"email" format:"email"`
	Provider string `json:"provider"`
	Status   string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connected   respjson.Field
		Email       respjson.Field
		Provider    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxCheckConnectionStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxCheckConnectionStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxConnectEmailResponse struct {
	// Seconds until the OAuth URL expires (typically 10 minutes)
	ExpiresIn int64 `json:"expires_in"`
	// Redirect user to this URL to start OAuth flow
	OAuthURL string `json:"oauth_url" format:"uri"`
	Status   string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresIn   respjson.Field
		OAuthURL    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxConnectEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxConnectEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxDisconnectEmailResponse struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxDisconnectEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxDisconnectEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxListCasFilesResponse struct {
	// Number of CAS files found
	Count  int64                           `json:"count"`
	Files  []InboxListCasFilesResponseFile `json:"files"`
	Status string                          `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Files       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxListCasFilesResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxListCasFilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A CAS file found in the user's email inbox
type InboxListCasFilesResponseFile struct {
	// Detected CAS provider based on sender email
	//
	// Any of "cdsl", "nsdl", "cams", "kfintech".
	CasType string `json:"cas_type"`
	// URL expiration time in seconds (default 86400 = 24 hours)
	ExpiresIn int64 `json:"expires_in"`
	// Standardized filename (provider_YYYYMMDD_uniqueid.pdf)
	Filename string `json:"filename"`
	// Date the email was received
	MessageDate time.Time `json:"message_date" format:"date"`
	// Unique identifier for the email message (use for subsequent API calls)
	MessageID string `json:"message_id"`
	// Original attachment filename from the email
	OriginalFilename string `json:"original_filename"`
	// Email address of the CAS authority who sent this
	SenderEmail string `json:"sender_email" format:"email"`
	// File size in bytes
	Size int64 `json:"size"`
	// Direct download URL (presigned, expires based on expires_in)
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CasType          respjson.Field
		ExpiresIn        respjson.Field
		Filename         respjson.Field
		MessageDate      respjson.Field
		MessageID        respjson.Field
		OriginalFilename respjson.Field
		SenderEmail      respjson.Field
		Size             respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxListCasFilesResponseFile) RawJSON() string { return r.JSON.raw }
func (r *InboxListCasFilesResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxCheckConnectionStatusParams struct {
	XInboxToken string `header:"x-inbox-token,required" json:"-"`
	paramObj
}

type InboxConnectEmailParams struct {
	// Your callback URL to receive the inbox_token (must be http or https)
	RedirectUri string `json:"redirect_uri,required" format:"uri"`
	// State parameter for CSRF protection (returned in redirect)
	State param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r InboxConnectEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxConnectEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxConnectEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxDisconnectEmailParams struct {
	XInboxToken string `header:"x-inbox-token,required" json:"-"`
	paramObj
}

type InboxListCasFilesParams struct {
	XInboxToken string `header:"x-inbox-token,required" json:"-"`
	// End date in ISO format (YYYY-MM-DD). Defaults to today.
	EndDate param.Opt[time.Time] `json:"end_date,omitzero" format:"date"`
	// Start date in ISO format (YYYY-MM-DD). Defaults to 30 days ago.
	StartDate param.Opt[time.Time] `json:"start_date,omitzero" format:"date"`
	// Filter by CAS provider(s):
	//
	// - `cdsl` → eCAS@cdslstatement.com
	// - `nsdl` → NSDL-CAS@nsdl.co.in
	// - `cams` → donotreply@camsonline.com
	// - `kfintech` → samfS@kfintech.com
	//
	// Any of "cdsl", "nsdl", "cams", "kfintech".
	CasTypes []string `json:"cas_types,omitzero"`
	paramObj
}

func (r InboxListCasFilesParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxListCasFilesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxListCasFilesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
