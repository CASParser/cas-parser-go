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
)

// Endpoints for parsing CAS PDF files from different sources.
//
// CdslService contains methods and other services that help with interacting with
// the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCdslService] method instead.
type CdslService struct {
	Options []option.RequestOption
	// Endpoints for fetching CAS documents with instant download. Currently supports
	// CDSL via OTP authentication.
	Fetch CdslFetchService
}

// NewCdslService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCdslService(opts ...option.RequestOption) (r CdslService) {
	r = CdslService{}
	r.Options = opts
	r.Fetch = NewCdslFetchService(opts...)
	return
}

// This endpoint specifically parses CDSL CAS (Consolidated Account Statement) PDF
// files and returns data in a unified format. Use this endpoint when you know the
// PDF is from CDSL.
func (r *CdslService) ParsePdf(ctx context.Context, body CdslParsePdfParams, opts ...option.RequestOption) (res *UnifiedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/cdsl/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CdslParsePdfParams struct {
	// Password for the PDF file (if required)
	Password param.Opt[string] `json:"password,omitzero"`
	// Base64 encoded CAS PDF file (required if pdf_url not provided)
	PdfFile param.Opt[string] `json:"pdf_file,omitzero" format:"base64"`
	// URL to the CAS PDF file (required if pdf_file not provided)
	PdfURL param.Opt[string] `json:"pdf_url,omitzero" format:"uri"`
	paramObj
}

func (r CdslParsePdfParams) MarshalJSON() (data []byte, err error) {
	type shadow CdslParsePdfParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CdslParsePdfParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
