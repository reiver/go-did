package syntax_test

import (
	"testing"

	"math/rand"
	"time"

	"github.com/reiver/go-arbitrary"

	"github.com/reiver/go-did/syntax"
)

func TestIsMethodName(t *testing.T) {

	tests := []struct{
		String string
		Expected bool
	}{
		{
			String: "",
			Expected: false,
		},



		{
			String: "fAil",
			Expected: false,
		},
		{
			String: "pass",
			Expected: true,
		},



		{
			String: "\x00",
			Expected: false,
		},
		{
			String: "\x01",
			Expected: false,
		},
		{
			String: "\x02",
			Expected: false,
		},
		{
			String: "\x03",
			Expected: false,
		},
		{
			String: "\x04",
			Expected: false,
		},
		{
			String: "\x05",
			Expected: false,
		},
		{
			String: "\x06",
			Expected: false,
		},
		{
			String: "\x07",
			Expected: false,
		},
		{
			String: "\x08",
			Expected: false,
		},
		{
			String: "\x09",
			Expected: false,
		},
		{
			String: "\x0A",
			Expected: false,
		},
		{
			String: "\x0B",
			Expected: false,
		},
		{
			String: "\x0C",
			Expected: false,
		},
		{
			String: "\x0D",
			Expected: false,
		},
		{
			String: "\x0E",
			Expected: false,
		},
		{
			String: "\x0F",
			Expected: false,
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
			String: "A",
			Expected: false,
		},
		{
			String: "B",
			Expected: false,
		},
		{
			String: "C",
			Expected: false,
		},
		{
			String: "D",
			Expected: false,
		},
		{
			String: "E",
			Expected: false,
		},
		{
			String: "F",
			Expected: false,
		},
		{
			String: "G",
			Expected: false,
		},
		{
			String: "H",
			Expected: false,
		},
		{
			String: "I",
			Expected: false,
		},
		{
			String: "J",
			Expected: false,
		},
		{
			String: "K",
			Expected: false,
		},
		{
			String: "L",
			Expected: false,
		},
		{
			String: "M",
			Expected: false,
		},
		{
			String: "N",
			Expected: false,
		},
		{
			String: "O",
			Expected: false,
		},
		{
			String: "P",
			Expected: false,
		},
		{
			String: "Q",
			Expected: false,
		},
		{
			String: "R",
			Expected: false,
		},
		{
			String: "S",
			Expected: false,
		},
		{
			String: "T",
			Expected: false,
		},
		{
			String: "U",
			Expected: false,
		},
		{
			String: "V",
			Expected: false,
		},
		{
			String: "W",
			Expected: false,
		},
		{
			String: "X",
			Expected: false,
		},
		{
			String: "Y",
			Expected: false,
		},
		{
			String: "Z",
			Expected: false,
		},
	}

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	{
		var chars string = "0123456789abcdefghijklmnopqrstuvwxyz"

		for i:=0; i<0xFF; i++ {

			var str string = arbitrary.String(2+randomness.Intn(30), chars)

			test := struct{
				String string
				Expected bool
			}{
				String:str,
				Expected:true,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actual := syntax.IsMethodName(test.String)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("STRING: %q", test.String)
			continue
		}
	}
}
