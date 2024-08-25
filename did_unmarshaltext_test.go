package did_test

import (
	"testing"

	"github.com/reiver/go-did"
)

func TestDID_UnmarshalText(t *testing.T) {

	tests := []struct{
		Bytes []byte
		Expected string
	}{
		{
			Bytes: []byte("did:example:123"),
			Expected:     "did:example:123",
		},



		{
			Bytes: []byte("did:apple:banana:cherry"),
			Expected:     "did:apple:banana:cherry",
		},
	}

	for testNumber, test := range tests {
		var actualDID did.DID

		err := actualDID.UnmarshalText(test.Bytes)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("BYTES: %q", string(test.Bytes))
			continue
		}

		{
			expected := test.Expected
			actual := actualDID.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("BYTES: %q", string(test.Bytes))
				continue
			}
		}
	}
}
