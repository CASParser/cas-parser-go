// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/cas-parser-go/internal/apijson"
	"github.com/stainless-sdks/cas-parser-go/internal/requestconfig"
	"github.com/stainless-sdks/cas-parser-go/option"
	"github.com/stainless-sdks/cas-parser-go/packages/param"
)

// SmartService contains methods and other services that help with interacting with
// the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmartService] method instead.
type SmartService struct {
	Options []option.RequestOption
}

// NewSmartService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSmartService(opts ...option.RequestOption) (r SmartService) {
	r = SmartService{}
	r.Options = opts
	return
}

// This endpoint parses CAS (Consolidated Account Statement) PDF files from NSDL,
// CDSL, or CAMS/KFintech and returns data in a unified format. It auto-detects the
// CAS type and transforms the data into a consistent structure regardless of the
// source.
func (r *SmartService) ParseCasPdf(ctx context.Context, body SmartParseCasPdfParams, opts ...option.RequestOption) (res *UnifiedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/smart/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SmartParseCasPdfParams struct {
	// Password for the PDF file (if required)
	Password param.Opt[string] `json:"password,omitzero"`
	// Base64 encoded CAS PDF file (required if pdf_url not provided)
	PdfFile param.Opt[string] `json:"pdf_file,omitzero" format:"base64"`
	// URL to the CAS PDF file (required if pdf_file not provided)
	PdfURL param.Opt[string] `json:"pdf_url,omitzero" format:"uri"`
	paramObj
}

func (r SmartParseCasPdfParams) MarshalJSON() (data []byte, err error) {
	type shadow SmartParseCasPdfParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmartParseCasPdfParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
