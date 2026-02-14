// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"github.com/CASParser/cas-parser-go/option"
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
