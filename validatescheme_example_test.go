package did_test

import (
	"fmt"

	"github.com/reiver/go-did"
)

func ExampleValidateScheme() {

	var uri string = "http://example.com/once/twice/thrice/fource.html"

	err := did.ValidateScheme(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: did: URI "http://example.com/once/twice/thrice/fource.html" is not a DID because it does not begin with "did:"
}
