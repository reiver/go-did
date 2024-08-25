package did_test

import (
	"testing"

	"github.com/reiver/go-did"
)

func TestDID_String(t *testing.T) {

	tests := []struct{
		DID did.DID
		Expected string
	}{
		{
			DID: did.DID("did:apple:123"),
			Expected:    "did:apple:123",
		},
		{
			DID: did.DID("did:banana:123"),
			Expected:    "did:banana:123",
		},
		{
			DID: did.DID("did:cherry:123"),
			Expected:    "did:cherry:123",
		},



		{
			DID: did.DID("did:apple:banana:cherry"),
			Expected:    "did:apple:banana:cherry",
		},



		{
			DID: did.DID("did"),
			Expected:    "",
		},
		{
			DID: did.DID("did:"),
			Expected:    "",
		},
		{
			DID: did.DID("abc:method:id"),
			Expected:    "",
		},
		{
			DID: did.DID("did:$$$:id"),
			Expected:    "",
		},
	}

	for testNumber, test := range tests {
		actual := test.DID.String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'string' value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}