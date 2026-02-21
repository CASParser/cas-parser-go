// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"github.com/CASParser/cas-parser-go/option"
)

// VerifyTokenService contains methods and other services that help with
// interacting with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerifyTokenService] method instead.
type VerifyTokenService struct {
	Options []option.RequestOption
}

// NewVerifyTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerifyTokenService(opts ...option.RequestOption) (r VerifyTokenService) {
	r = VerifyTokenService{}
	r.Options = opts
	return
}
