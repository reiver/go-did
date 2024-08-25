package syntax

// IsMethodName returns true if `value` is a method-name, as defined by https://www.w3.org/TR/did-core/
//
// Which is defined as:
//
//	method-name = 1*method-char
//	
//	method-char = %x61-7A / DIGIT
//	
//	DIGIT       = %x30-39 ; 0-9
func IsMethodName(value string) bool {
	var length int = len(value)

	if length < 1 {
		return false
	}

	for _, r := range value {

		if !IsMethodChar(r) {
			return false
		}
	}

	return true
}
