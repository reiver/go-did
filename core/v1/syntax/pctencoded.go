package syntax

import (
	"github.com/reiver/go-rfc3986"
)

// HasPrefixPctEncoded returns true if `str` starts with pct-encoded, as defined by https://www.w3.org/TR/did-core/
//
// Which is defined as:
//
//	pct-encoded = "%" HEXDIG HEXDIG
//	
//	HEXDIG      = DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
//
// Note that `HEXDIG` does NOT allow lower-case letters "a", "b", "c", "d", "e", "f".
// Only upper-case letters "A", "B", "C", "D", "E", "F".
func HasPrefixPctEncoded(str string) bool {
	return rfc3986.HasPrefixPctEncoded(str)
}
