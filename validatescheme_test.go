package did_test

import (
	"testing"

	"github.com/reiver/go-did"
)

func TestValidateScheme(t *testing.T) {

	tests := []struct{
		URI string
	}{
		{
			URI: "did:",
		},
		{
			URI: "diD:",
		},
		{
			URI: "dId:",
		},
		{
			URI: "dID:",
		},
		{
			URI: "Did:",
		},
		{
			URI: "DiD:",
		},
		{
			URI: "DId:",
		},
		{
			URI: "DID:",
		},



		{
			URI: "did://",
		},
		{
			URI: "diD://",
		},
		{
			URI: "dId://",
		},
		{
			URI: "dID://",
		},
		{
			URI: "Did://",
		},
		{
			URI: "DiD://",
		},
		{
			URI: "DId://",
		},
		{
			URI: "DID://",
		},



		{
			URI: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "diD:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "dId:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "dID:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "Did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "DiD:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "DId:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "DID:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI: "did://VIDEO.Archive.ORG",
		},
		{
			URI: "diD://VIDEO.Archive.ORG",
		},
		{
			URI: "dId://VIDEO.Archive.ORG",
		},
		{
			URI: "dID://VIDEO.Archive.ORG",
		},
		{
			URI: "Did://VIDEO.Archive.ORG",
		},
		{
			URI: "DiD://VIDEO.Archive.ORG",
		},
		{
			URI: "DId://VIDEO.Archive.ORG",
		},
		{
			URI: "DID://VIDEO.Archive.ORG",
		},



		{
			URI: "DID://user:pass@SOMETHING.Example/banana/wxyz",
		},
	}

	for testNumber, test := range tests {

		err := did.ValidateScheme(test.URI)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}

func TestValidateScheme_fail(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: `did: empty URI`,
		},



		{
			URI: "d",
			ExpectedError: `did: URI "d" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "D",
			ExpectedError: `did: URI "D" is not a DID because it does not begin with "did:"`,
		},



		{
			URI: "di",
			ExpectedError: `did: URI "di" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "dI",
			ExpectedError: `did: URI "dI" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "Di",
			ExpectedError: `did: URI "Di" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "DI",
			ExpectedError: `did: URI "DI" is not a DID because it does not begin with "did:"`,
		},



		{
			URI: "did",
			ExpectedError: `did: URI "did" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "diD",
			ExpectedError: `did: URI "diD" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "dId",
			ExpectedError: `did: URI "dId" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "dID",
			ExpectedError: `did: URI "dID" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "Did",
			ExpectedError: `did: URI "Did" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "DiD",
			ExpectedError: `did: URI "DiD" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "DId",
			ExpectedError: `did: URI "DId" is not a DID because it does not begin with "did:"`,
		},
		{
			URI: "DID",
			ExpectedError: `did: URI "DID" is not a DID because it does not begin with "did:"`,
		},



		{
			URI: "http://example.com",
			ExpectedError: `did: URI "http://example.com" is not a DID because it does not begin with "did:"`,
		},
	}

	for testNumber, test := range tests {

		err := did.ValidateScheme(test.URI)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("URI: %s", test.URI)
			continue
		}

		actual := err.Error()
		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED-ERROR: %s", expected)
			t.Logf("ACTUAL-ERROR:   %s", actual)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}
