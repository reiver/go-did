package syntax

// IsLCAlpha returns true if `value` is a LCALPHA, as (implicitly) defined by https://www.w3.org/TR/did-core/
//
// (It is implicitly defined as part of 'method-char'.)
//
// Which is defined as:
//
//	LCALPHA = %x61-7A ; a-z
func IsLCAlpha(value rune) bool {
	return 'a' <= value && value <= 'z'
}
