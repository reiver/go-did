package did

import (
	"strings"

	"github.com/reiver/go-erorr"
)

const schemePrefix string = "did:"
const lenSchemePrefix int = len(schemePrefix)

// ValidateScheme only validates the scheme of a potential DID.
//
// So, it checks to see if the URI starts with "did:".
// And, that is it.
//
// You would use ValidateScheme if you wanted to be very liberal in what you accept as a valid DID, including not caring if the DID has a 'method-name' or a 'method-specific-id'.
// I.e., this is the minimum amount of validation you can do to validate a DID.
func ValidateScheme(uri string) error {

	if "" == uri {
		return errEmptyURI
	}

	{
		var lenuri int = len(uri)
		if lenuri < lenSchemePrefix {
			return erorr.Errorf("did: URI %q is not a DID because it does not begin with %q", uri, schemePrefix)
		}

                var beginning string = uri[:lenSchemePrefix]

                if strings.ToLower(beginning) != schemePrefix {
			return erorr.Errorf("did: URI %q is not a DID because it does not begin with %q", uri, schemePrefix)
                }
	}

	return nil
}
