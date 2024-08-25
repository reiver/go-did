package did_test

import (
	"testing"

	"github.com/reiver/go-did"
)

func TestDID_Identifier(t *testing.T) {

	tests := []struct{
		DID did.DID
		Expected string
	}{
		{
			DID: did.DID("did:apple:123"),
			Expected:              "123",
		},
		{
			DID: did.DID("did:banana:123"),
			Expected:               "123",
		},
		{
			DID: did.DID("did:cherry:123"),
			Expected:               "123",
		},



		{
			DID: did.DID("did:apple:banana:cherry"),
			Expected:              "banana:cherry",
		},



		{
			DID: did.DID("did:example123:0::123456789-ABCDEFGHIJKLMNOPQRSTUXWXYZ_:."),
			Expected:                   "0::123456789-ABCDEFGHIJKLMNOPQRSTUXWXYZ_:.",
		},



		{
			DID: did.DID("did"),
			Expected: "",
		},
		{
			DID: did.DID("did:"),
			Expected: "",
		},
		{
			DID: did.DID("abc:method:id"),
			Expected: "",
		},
		{
			DID: did.DID("did:$$$:id"),
			Expected: "",
		},
	}

	for testNumber, test := range tests {
		actual := test.DID.Identifier()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'identifier' value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
