// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/CASParser/cas-parser-go/internal/apijson"
	"github.com/CASParser/cas-parser-go/internal/apiquery"
	"github.com/CASParser/cas-parser-go/internal/requestconfig"
	"github.com/CASParser/cas-parser-go/option"
	"github.com/CASParser/cas-parser-go/packages/param"
	"github.com/CASParser/cas-parser-go/packages/respjson"
)

// Create dedicated inbound email addresses for investors to forward their CAS
// statements.
//
// **Use Case:** Your app wants to collect CAS statements from users without
// requiring OAuth or file upload.
//
// **How it works:**
//
//  1. Call `POST /v4/inbound-email` to create a unique inbound email address
//  2. Display this email to your user: "Forward your CAS statement to
//     ie_xxx@import.casparser.in"
//  3. When user forwards a CAS email, we verify sender authenticity (SPF/DKIM) and
//     call your webhook
//  4. Your webhook receives email metadata + attachment download URLs
//
// **Sender Validation:**
//
// - Only emails from verified CAS authorities are processed:
//   - CDSL: `eCAS@cdslstatement.com`
//   - NSDL: `NSDL-CAS@nsdl.co.in`
//   - CAMS: `donotreply@camsonline.com`
//   - KFintech: `samfS@kfintech.com`
//
// - Emails failing SPF/DKIM/DMARC are rejected
// - Forwarded emails must contain the original sender in headers
//
// **Billing:** 0.2 credits per successfully processed valid email
//
// InboundEmailService contains methods and other services that help with
// interacting with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundEmailService] method instead.
type InboundEmailService struct {
	Options []option.RequestOption
}

// NewInboundEmailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboundEmailService(opts ...option.RequestOption) (r InboundEmailService) {
	r = InboundEmailService{}
	r.Options = opts
	return
}

// Create a dedicated inbound email address for collecting CAS statements via email
// forwarding. When an investor forwards a CAS email to this address, we verify the
// sender and make the file available to you.
//
// `callback_url` is **optional**:
//
//   - **Set it** — we POST each parsed email to your webhook as it arrives.
//   - **Omit it** — retrieve files via `GET /v4/inbound-email/{id}/files` without
//     building a webhook consumer.
func (r *InboundEmailService) New(ctx context.Context, body InboundEmailNewParams, opts ...option.RequestOption) (res *InboundEmailNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/inbound-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details of a specific inbound email including statistics.
func (r *InboundEmailService) Get(ctx context.Context, inboundEmailID string, opts ...option.RequestOption) (res *InboundEmailGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundEmailID == "" {
		err = errors.New("missing required inbound_email_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v4/inbound-email/%s", url.PathEscape(inboundEmailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all inbound emails associated with your API key. Returns active and paused
// inbound emails (deleted ones are excluded).
func (r *InboundEmailService) List(ctx context.Context, query InboundEmailListParams, opts ...option.RequestOption) (res *InboundEmailListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/inbound-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently delete an inbound email address. It will stop accepting emails.
//
// **Note:** Deletion is immediate and cannot be undone. Any emails received after
// deletion will be rejected.
func (r *InboundEmailService) Delete(ctx context.Context, inboundEmailID string, opts ...option.RequestOption) (res *InboundEmailDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundEmailID == "" {
		err = errors.New("missing required inbound_email_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v4/inbound-email/%s", url.PathEscape(inboundEmailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// An inbound email address for receiving forwarded CAS emails
type InboundEmailNewResponse struct {
	// Accepted CAS providers (empty = all)
	//
	// Any of "cdsl", "nsdl", "cams", "kfintech".
	AllowedSources []string `json:"allowed_sources"`
	// Webhook URL for email notifications. If set, we POST each parsed email here. If
	// omitted, files are only retrievable via `GET /v4/inbound-email/{id}/files`.
	CallbackURL string `json:"callback_url" format:"uri"`
	// When the inbound email was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The inbound email address to forward CAS statements to
	Email string `json:"email" format:"email"`
	// Unique inbound email identifier
	InboundEmailID string `json:"inbound_email_id"`
	// Custom key-value metadata
	Metadata map[string]string `json:"metadata"`
	// Your internal reference identifier
	Reference string `json:"reference" api:"nullable"`
	// Current inbound email lifecycle status
	//
	// Any of "active", "paused".
	Status InboundEmailNewResponseStatus `json:"status"`
	// When the inbound email was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedSources respjson.Field
		CallbackURL    respjson.Field
		CreatedAt      respjson.Field
		Email          respjson.Field
		InboundEmailID respjson.Field
		Metadata       respjson.Field
		Reference      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundEmailNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundEmailNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current inbound email lifecycle status
type InboundEmailNewResponseStatus string

const (
	InboundEmailNewResponseStatusActive InboundEmailNewResponseStatus = "active"
	InboundEmailNewResponseStatusPaused InboundEmailNewResponseStatus = "paused"
)

// An inbound email address for receiving forwarded CAS emails
type InboundEmailGetResponse struct {
	// Accepted CAS providers (empty = all)
	//
	// Any of "cdsl", "nsdl", "cams", "kfintech".
	AllowedSources []string `json:"allowed_sources"`
	// Webhook URL for email notifications. If set, we POST each parsed email here. If
	// omitted, files are only retrievable via `GET /v4/inbound-email/{id}/files`.
	CallbackURL string `json:"callback_url" format:"uri"`
	// When the inbound email was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The inbound email address to forward CAS statements to
	Email string `json:"email" format:"email"`
	// Unique inbound email identifier
	InboundEmailID string `json:"inbound_email_id"`
	// Custom key-value metadata
	Metadata map[string]string `json:"metadata"`
	// Your internal reference identifier
	Reference string `json:"reference" api:"nullable"`
	// Current inbound email lifecycle status
	//
	// Any of "active", "paused".
	Status InboundEmailGetResponseStatus `json:"status"`
	// When the inbound email was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedSources respjson.Field
		CallbackURL    respjson.Field
		CreatedAt      respjson.Field
		Email          respjson.Field
		InboundEmailID respjson.Field
		Metadata       respjson.Field
		Reference      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundEmailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundEmailGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current inbound email lifecycle status
type InboundEmailGetResponseStatus string

const (
	InboundEmailGetResponseStatusActive InboundEmailGetResponseStatus = "active"
	InboundEmailGetResponseStatusPaused InboundEmailGetResponseStatus = "paused"
)

type InboundEmailListResponse struct {
	InboundEmails []InboundEmailListResponseInboundEmail `json:"inbound_emails"`
	Limit         int64                                  `json:"limit"`
	Offset        int64                                  `json:"offset"`
	Status        string                                 `json:"status"`
	// Total number of inbound emails (for pagination)
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundEmails respjson.Field
		Limit         respjson.Field
		Offset        respjson.Field
		Status        respjson.Field
		Total         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundEmailListResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundEmailListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An inbound email address for receiving forwarded CAS emails
type InboundEmailListResponseInboundEmail struct {
	// Accepted CAS providers (empty = all)
	//
	// Any of "cdsl", "nsdl", "cams", "kfintech".
	AllowedSources []string `json:"allowed_sources"`
	// Webhook URL for email notifications. If set, we POST each parsed email here. If
	// omitted, files are only retrievable via `GET /v4/inbound-email/{id}/files`.
	CallbackURL string `json:"callback_url" format:"uri"`
	// When the inbound email was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The inbound email address to forward CAS statements to
	Email string `json:"email" format:"email"`
	// Unique inbound email identifier
	InboundEmailID string `json:"inbound_email_id"`
	// Custom key-value metadata
	Metadata map[string]string `json:"metadata"`
	// Your internal reference identifier
	Reference string `json:"reference" api:"nullable"`
	// Current inbound email lifecycle status
	//
	// Any of "active", "paused".
	Status string `json:"status"`
	// When the inbound email was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedSources respjson.Field
		CallbackURL    respjson.Field
		CreatedAt      respjson.Field
		Email          respjson.Field
		InboundEmailID respjson.Field
		Metadata       respjson.Field
		Reference      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundEmailListResponseInboundEmail) RawJSON() string { return r.JSON.raw }
func (r *InboundEmailListResponseInboundEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundEmailDeleteResponse struct {
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
func (r InboundEmailDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundEmailDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundEmailNewParams struct {
	// Optional webhook URL where we POST parsed emails. Must be HTTPS in production
	// (HTTP allowed for localhost). If omitted, retrieve files via
	// `GET /v4/inbound-email/{id}/files`.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero" format:"uri"`
	// Optional custom email prefix (e.g. `john-portfolio@import.casparser.in`). 3-32
	// chars, alphanumeric + hyphens, must start/end with a letter or number. If
	// omitted, a random ID is generated.
	Alias param.Opt[string] `json:"alias,omitzero"`
	// Your internal identifier (e.g., user_id, account_id). Returned in webhook
	// payload for correlation.
	Reference param.Opt[string] `json:"reference,omitzero"`
	// Filter emails by CAS provider. If omitted, accepts all providers.
	//
	// - `cdsl` → eCAS@cdslstatement.com
	// - `nsdl` → NSDL-CAS@nsdl.co.in
	// - `cams` → donotreply@camsonline.com
	// - `kfintech` → samfS@kfintech.com
	//
	// Any of "cdsl", "nsdl", "cams", "kfintech".
	AllowedSources []string `json:"allowed_sources,omitzero"`
	// Optional key-value pairs (max 10) to include in webhook payload. Useful for
	// passing context like plan_type, campaign_id, etc.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r InboundEmailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InboundEmailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundEmailNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundEmailListParams struct {
	// Maximum number of inbound emails to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by status
	//
	// Any of "active", "paused", "all".
	Status InboundEmailListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboundEmailListParams]'s query parameters as `url.Values`.
func (r InboundEmailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status
type InboundEmailListParamsStatus string

const (
	InboundEmailListParamsStatusActive InboundEmailListParamsStatus = "active"
	InboundEmailListParamsStatusPaused InboundEmailListParamsStatus = "paused"
	InboundEmailListParamsStatusAll    InboundEmailListParamsStatus = "all"
)
