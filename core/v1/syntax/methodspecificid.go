package syntax

// IsMethodSpecificID returns true if `value` is a method-specific-id, as defined by https://www.w3.org/TR/did-core/
//
// Which is defined as:
//
//	method-specific-id = *( *idchar ":" ) 1*idchar
//	
//	idchar             = ALPHA / DIGIT / "." / "-" / "_" / pct-encoded
//	
//	ALPHA              = %x41-5A / %x61-7A ; A-Z / a-z
//	
//	DIGIT              = %x30-39 ; 0-9
//	
//	pct-encoded        = "%" HEXDIG HEXDIG
//	
//	HEXDIG         =  DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
func IsMethodSpecificID(value string) bool {
	if len(value) < 1 {
		return false
	}

	var last rune

	for i, r := range value {
		var s string = value[i:]

		if !HasPrefixIDChar(s) && ':' != r {
			return false
		}

		last = r
	}

	if ':' == last {
		return false
	}

	return true
}
