package did_test

import (
	"testing"

	"github.com/reiver/go-did"
)

func TestMustConstructDID(t *testing.T) {

	tests := []struct{
		MethodName string
		Identifier string
		Expected string
	}{
		{
			MethodName:   "apple",
			Identifier:         "123",
			Expected: "did:apple:123",
		},
		{
			MethodName:   "banana",
			Identifier:          "456:7:::89-_.",
			Expected: "did:banana:456:7:::89-_.",
		},
		{
			MethodName:   "cherry",
			Identifier:          "juice.txt",
			Expected: "did:cherry:juice.txt",
		},



		{
			MethodName:   "0",
			Identifier:     ":-",
			Expected: "did:0::-",
		},
	}

	for testNumber, test := range tests {

		actualDID := did.MustConstructDID(test.MethodName, test.Identifier)
		actual := actualDID.String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'did' it not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("METHOD-NAME: %q", test.MethodName)
			t.Logf("IDENTIFIER:  %q", test.Identifier)
			continue
		}
	}
}
