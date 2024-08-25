package syntax_test

import (
	"testing"

	"github.com/reiver/go-did/core/v1/syntax"
)

func TestIsMethodSpecificID(t *testing.T) {

	tests := []struct{
		String string
		Expected bool
	}{
		{
			String: "",
			Expected: false,
		},



		{
			String: ":",
			Expected: false,
		},



		{
			String: "apple:",
			Expected: false,
		},
		{
			String: "banana:",
			Expected: false,
		},
		{
			String: "cherry:",
			Expected: false,
		},



		{
			String: ":0",
			Expected: true,
		},
		{
			String: ":1",
			Expected: true,
		},
		{
			String: ":2",
			Expected: true,
		},
		{
			String: ":3",
			Expected: true,
		},
		{
			String: ":4",
			Expected: true,
		},
		{
			String: ":5",
			Expected: true,
		},
		{
			String: ":6",
			Expected: true,
		},
		{
			String: ":7",
			Expected: true,
		},
		{
			String: ":8",
			Expected: true,
		},
		{
			String: ":9",
			Expected: true,
		},



		{
			String: "0",
			Expected: true,
		},
		{
			String: "1",
			Expected: true,
		},
		{
			String: "2",
			Expected: true,
		},
		{
			String: "3",
			Expected: true,
		},
		{
			String: "4",
			Expected: true,
		},
		{
			String: "5",
			Expected: true,
		},
		{
			String: "6",
			Expected: true,
		},
		{
			String: "7",
			Expected: true,
		},
		{
			String: "8",
			Expected: true,
		},
		{
			String: "9",
			Expected: true,
		},



		{
			String: "A",
			Expected: true,
		},
		{
			String: "B",
			Expected: true,
		},
		{
			String: "C",
			Expected: true,
		},
		{
			String: "D",
			Expected: true,
		},
		{
			String: "E",
			Expected: true,
		},
		{
			String: "F",
			Expected: true,
		},
		{
			String: "G",
			Expected: true,
		},
		{
			String: "H",
			Expected: true,
		},
		{
			String: "I",
			Expected: true,
		},
		{
			String: "J",
			Expected: true,
		},
		{
			String: "K",
			Expected: true,
		},
		{
			String: "L",
			Expected: true,
		},
		{
			String: "M",
			Expected: true,
		},
		{
			String: "N",
			Expected: true,
		},
		{
			String: "O",
			Expected: true,
		},
		{
			String: "P",
			Expected: true,
		},
		{
			String: "Q",
			Expected: true,
		},
		{
			String: "R",
			Expected: true,
		},
		{
			String: "S",
			Expected: true,
		},
		{
			String: "T",
			Expected: true,
		},
		{
			String: "U",
			Expected: true,
		},
		{
			String: "V",
			Expected: true,
		},
		{
			String: "W",
			Expected: true,
		},
		{
			String: "X",
			Expected: true,
		},
		{
			String: "Y",
			Expected: true,
		},
		{
			String: "Z",
			Expected: true,
		},



		{
			String: "a",
			Expected: true,
		},
		{
			String: "b",
			Expected: true,
		},
		{
			String: "c",
			Expected: true,
		},
		{
			String: "d",
			Expected: true,
		},
		{
			String: "e",
			Expected: true,
		},
		{
			String: "f",
			Expected: true,
		},
		{
			String: "g",
			Expected: true,
		},
		{
			String: "h",
			Expected: true,
		},
		{
			String: "i",
			Expected: true,
		},
		{
			String: "j",
			Expected: true,
		},
		{
			String: "k",
			Expected: true,
		},
		{
			String: "l",
			Expected: true,
		},
		{
			String: "m",
			Expected: true,
		},
		{
			String: "n",
			Expected: true,
		},
		{
			String: "o",
			Expected: true,
		},
		{
			String: "p",
			Expected: true,
		},
		{
			String: "q",
			Expected: true,
		},
		{
			String: "r",
			Expected: true,
		},
		{
			String: "s",
			Expected: true,
		},
		{
			String: "t",
			Expected: true,
		},
		{
			String: "u",
			Expected: true,
		},
		{
			String: "v",
			Expected: true,
		},
		{
			String: "w",
			Expected: true,
		},
		{
			String: "x",
			Expected: true,
		},
		{
			String: "y",
			Expected: true,
		},
		{
			String: "z",
			Expected: true,
		},



		{
			String: ".",
			Expected: true,
		},
		{
			String: "-",
			Expected: true,
		},
		{
			String: "_",
			Expected: true,
		},



		{
			String: "%00",
			Expected: true,
		},
		{
			String: "%01",
			Expected: true,
		},
		{
			String: "%02",
			Expected: true,
		},
		{
			String: "%03",
			Expected: true,
		},
		{
			String: "%04",
			Expected: true,
		},
		{
			String: "%05",
			Expected: true,
		},
		{
			String: "%06",
			Expected: true,
		},
		{
			String: "%07",
			Expected: true,
		},
		{
			String: "%08",
			Expected: true,
		},
		{
			String: "%09",
			Expected: true,
		},
		{
			String: "%0A",
			Expected: true,
		},
		{
			String: "%0B",
			Expected: true,
		},
		{
			String: "%0C",
			Expected: true,
		},
		{
			String: "%0D",
			Expected: true,
		},
		{
			String: "%0E",
			Expected: true,
		},
		{
			String: "%0F",
			Expected: true,
		},



		{
			String: "%FF",
			Expected: true,
		},
	}

	for testNumber, test := range tests {
		actual := syntax.IsMethodSpecificID(test.String)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL: %t", actual)
			t.Logf("STRING: %q", test.String)
			continue
		}
	}
}
