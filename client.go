// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/CASParser/cas-parser-go/internal/requestconfig"
	"github.com/CASParser/cas-parser-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cas-parser API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	// Endpoints for checking API quota and credits usage. These endpoints help you
	// monitor your API usage and remaining quota.
	Credits CreditService
	// Endpoints for checking API quota and credits usage. These endpoints help you
	// monitor your API usage and remaining quota.
	Logs LogService
	// Endpoints for managing access tokens for the Portfolio Connect SDK. Use these to
	// generate short-lived `at_` prefixed tokens that can be safely passed to frontend
	// applications. Access tokens can be used in place of API keys on all v4
	// endpoints.
	AccessToken AccessTokenService
	// Endpoints for managing access tokens for the Portfolio Connect SDK. Use these to
	// generate short-lived `at_` prefixed tokens that can be safely passed to frontend
	// applications. Access tokens can be used in place of API keys on all v4
	// endpoints.
	VerifyToken VerifyTokenService
	// Endpoints for parsing CAS PDF files from different sources.
	CamsKfintech CamsKfintechService
	// Endpoints for parsing CAS PDF files from different sources.
	Cdsl CdslService
	// Endpoints for parsing Contract Note PDF files from various SEBI brokers like
	// Zerodha, Groww, Upstox, ICICI etc.
	ContractNote ContractNoteService
	// Endpoints for importing CAS files directly from user email inboxes.
	//
	// **Supported Providers:** Gmail (more coming soon)
	//
	// **How it works:**
	//
	//  1. Call `POST /v4/inbox/connect` to get an OAuth URL
	//  2. Redirect user to the OAuth URL for consent
	//  3. User is redirected back to your `redirect_uri` with an encrypted
	//     `inbox_token`
	//  4. Use the token to list/fetch CAS files from their inbox (`/v4/inbox/cas`)
	//  5. Files are uploaded to temporary cloud storage (URLs expire in 24 hours)
	//
	// **Security:**
	//
	// - Read-only access (we cannot send emails)
	// - Tokens are encrypted with server-side secret
	// - User can revoke access anytime via `/v4/inbox/disconnect`
	Inbox InboxService
	// Endpoints for generating new CAS documents via email mailback (KFintech).
	Kfintech KfintechService
	// Endpoints for parsing CAS PDF files from different sources.
	Nsdl NsdlService
	// Endpoints for parsing CAS PDF files from different sources.
	Smart SmartService
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
	InboundEmail InboundEmailService
}

// DefaultClientOptions read from the environment (CAS_PARSER_API_KEY,
// CAS_PARSER_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CAS_PARSER_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("CAS_PARSER_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (CAS_PARSER_API_KEY, CAS_PARSER_BASE_URL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Credits = NewCreditService(opts...)
	r.Logs = NewLogService(opts...)
	r.AccessToken = NewAccessTokenService(opts...)
	r.VerifyToken = NewVerifyTokenService(opts...)
	r.CamsKfintech = NewCamsKfintechService(opts...)
	r.Cdsl = NewCdslService(opts...)
	r.ContractNote = NewContractNoteService(opts...)
	r.Inbox = NewInboxService(opts...)
	r.Kfintech = NewKfintechService(opts...)
	r.Nsdl = NewNsdlService(opts...)
	r.Smart = NewSmartService(opts...)
	r.InboundEmail = NewInboundEmailService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
