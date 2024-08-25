package syntax

// IsMethodChar returns true if `value` is a idchar, as defined by https://www.w3.org/TR/did-core/
//
// Which is defined as:
//
//	method-char = %x61-7A / DIGIT
//	
//	DIGIT       = %x30-39 ; 0-9
func IsMethodChar(value rune) bool {
	return IsLCAlpha(value) || IsDigit(value)
}
