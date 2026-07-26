// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"github.com/CASParser/cas-parser-go/option"
)

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
