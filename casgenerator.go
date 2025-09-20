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

// CasGeneratorService contains methods and other services that help with
// interacting with the CAS Parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasGeneratorService] method instead.
type CasGeneratorService struct {
	Options []option.RequestOption
}

// NewCasGeneratorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCasGeneratorService(opts ...option.RequestOption) (r CasGeneratorService) {
	r = CasGeneratorService{}
	r.Options = opts
	return
}

// This endpoint generates CAS (Consolidated Account Statement) documents by
// submitting a mailback request to the specified CAS authority. Currently only
// supports KFintech, with plans to support CAMS, CDSL, and NSDL in the future.
func (r *CasGeneratorService) GenerateCas(ctx context.Context, body CasGeneratorGenerateCasParams, opts ...option.RequestOption) (res *CasGeneratorGenerateCasResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CasGeneratorGenerateCasResponse struct {
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
func (r CasGeneratorGenerateCasResponse) RawJSON() string { return r.JSON.raw }
func (r *CasGeneratorGenerateCasResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CasGeneratorGenerateCasParams struct {
	// Email address to receive the CAS document
	Email string `json:"email,required"`
	// Start date for the CAS period (format YYYY-MM-DD)
	FromDate string `json:"from_date,required"`
	// Password to protect the generated CAS PDF
	Password string `json:"password,required"`
	// End date for the CAS period (format YYYY-MM-DD)
	ToDate string `json:"to_date,required"`
	// PAN number (optional for some CAS authorities)
	PanNo param.Opt[string] `json:"pan_no,omitzero"`
	// CAS authority to generate the document from (currently only kfintech is
	// supported)
	//
	// Any of "kfintech", "cams", "cdsl", "nsdl".
	CasAuthority CasGeneratorGenerateCasParamsCasAuthority `json:"cas_authority,omitzero"`
	paramObj
}

func (r CasGeneratorGenerateCasParams) MarshalJSON() (data []byte, err error) {
	type shadow CasGeneratorGenerateCasParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CasGeneratorGenerateCasParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CAS authority to generate the document from (currently only kfintech is
// supported)
type CasGeneratorGenerateCasParamsCasAuthority string

const (
	CasGeneratorGenerateCasParamsCasAuthorityKfintech CasGeneratorGenerateCasParamsCasAuthority = "kfintech"
	CasGeneratorGenerateCasParamsCasAuthorityCams     CasGeneratorGenerateCasParamsCasAuthority = "cams"
	CasGeneratorGenerateCasParamsCasAuthorityCdsl     CasGeneratorGenerateCasParamsCasAuthority = "cdsl"
	CasGeneratorGenerateCasParamsCasAuthorityNsdl     CasGeneratorGenerateCasParamsCasAuthority = "nsdl"
)
