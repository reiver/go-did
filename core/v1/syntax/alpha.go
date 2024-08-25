package syntax

import (
	"github.com/reiver/go-rfc2234"
)

// IsAlpha returns true if `r` is ALPHA, as defined by https://www.w3.org/TR/did-core/
//
// Which is defined as:
//
//	ALPHA = %x41-5A / %x61-7A ; A-Z / a-z
func IsAlpha(r rune) bool {
	return rfc2234.IsAlpha(r)
}
