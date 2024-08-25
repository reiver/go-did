package syntax

import (
	"github.com/reiver/go-utf8"
)

// HasPrefixIDChar returns true if `str` starts with idchar, as defined by https://www.w3.org/TR/did-core/
//
// Which is defined as:
//
//	idchar      = ALPHA / DIGIT / "." / "-" / "_" / pct-encoded
//	
//	pct-encoded = "%" HEXDIG HEXDIG
//	
//	HEXDIG      = DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
//
// Note that `HEXDIG` does NOT allow lower-case letters "a", "b", "c", "d", "e", "f".
// Only upper-case letters "A", "B", "C", "D", "E", "F".
func HasPrefixIDChar(value string) bool {
	var length int = len(value)

	if length < 1  {
		return false
	}

	r, _, err := utf8.DecodeFirstRuneOfString(value)
	if nil != err {
		return false
	}

	switch {
	case IsAlpha(r):
		return true
	case IsDigit(r):
		return true
	case '.' == r:
		return true
	case '-' == r:
		return true
	case '_' == r:
		return true
	case '%' == r:
		if !HasPrefixPctEncoded(value) {
			return false
		}
		return true
	default:
		return false
	}
}
