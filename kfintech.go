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

// Endpoints for generating new CAS documents via email mailback (KFintech).
//
// KfintechService contains methods and other services that help with interacting
// with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKfintechService] method instead.
type KfintechService struct {
	Options []option.RequestOption
}

// NewKfintechService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKfintechService(opts ...option.RequestOption) (r KfintechService) {
	r = KfintechService{}
	r.Options = opts
	return
}

// Generate CAS via KFintech mailback. The CAS PDF will be sent to the investor's
// email.
//
// This is an async operation - the investor receives the CAS via email within a
// few minutes. For instant CAS retrieval, use CDSL Fetch (`/v4/cdsl/fetch`).
func (r *KfintechService) GenerateCas(ctx context.Context, body KfintechGenerateCasParams, opts ...option.RequestOption) (res *KfintechGenerateCasResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/kfintech/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type KfintechGenerateCasResponse struct {
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
func (r KfintechGenerateCasResponse) RawJSON() string { return r.JSON.raw }
func (r *KfintechGenerateCasResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KfintechGenerateCasParams struct {
	// Email address to receive the CAS document
	Email string `json:"email" api:"required"`
	// Start date (YYYY-MM-DD)
	FromDate string `json:"from_date" api:"required"`
	// Password for the PDF
	Password string `json:"password" api:"required"`
	// End date (YYYY-MM-DD)
	ToDate string `json:"to_date" api:"required"`
	// PAN number (optional)
	PanNo param.Opt[string] `json:"pan_no,omitzero"`
	paramObj
}

func (r KfintechGenerateCasParams) MarshalJSON() (data []byte, err error) {
	type shadow KfintechGenerateCasParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KfintechGenerateCasParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
