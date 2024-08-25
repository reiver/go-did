package syntax

import (
	"github.com/reiver/go-rfc3986"
)

// IsDigit returns true if `r` is DIGIT, as defined by https://www.w3.org/TR/did-core/
//
// Which is defined as:
//
//	DIGIT = %x30-39 ; 0-9
func IsDigit(r rune) bool {
	return rfc3986.IsDigit(r)
}
