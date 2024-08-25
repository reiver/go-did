package syntax_test

import (
	"testing"

	"math/rand"
	"time"

	"github.com/reiver/go-arbitrary"

	"github.com/reiver/go-did/syntax"
)

func TestHasPrefixIDChar(t *testing.T) {

	tests := []struct{
		String string
		Expected bool
	}{
		{
			String: "",
			Expected: false,
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
			String: "%",
			Expected: false,
		},



		{
			String: "%0",
			Expected: false,
		},
		{
			String: "%1",
			Expected: false,
		},
		{
			String: "%2",
			Expected: false,
		},
		{
			String: "%3",
			Expected: false,
		},
		{
			String: "%4",
			Expected: false,
		},
		{
			String: "%5",
			Expected: false,
		},
		{
			String: "%6",
			Expected: false,
		},
		{
			String: "%7",
			Expected: false,
		},
		{
			String: "%8",
			Expected: false,
		},
		{
			String: "%9",
			Expected: false,
		},
		{
			String: "%A",
			Expected: false,
		},
		{
			String: "%B",
			Expected: false,
		},
		{
			String: "%C",
			Expected: false,
		},
		{
			String: "%D",
			Expected: false,
		},
		{
			String: "%E",
			Expected: false,
		},
		{
			String: "%F",
			Expected: false,
		},



		{
			String: "%a",
			Expected: false,
		},
		{
			String: "%b",
			Expected: false,
		},
		{
			String: "%c",
			Expected: false,
		},
		{
			String: "%d",
			Expected: false,
		},
		{
			String: "%e",
			Expected: false,
		},
		{
			String: "%f",
			Expected: false,
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
			String: "%0a",
			Expected: false,
		},
		{
			String: "%0b",
			Expected: false,
		},
		{
			String: "%0c",
			Expected: false,
		},
		{
			String: "%0d",
			Expected: false,
		},
		{
			String: "%0e",
			Expected: false,
		},
		{
			String: "%0f",
			Expected: false,
		},



		{
			String: "%10",
			Expected: true,
		},
		{
			String: "%11",
			Expected: true,
		},
		{
			String: "%12",
			Expected: true,
		},
		{
			String: "%13",
			Expected: true,
		},
		{
			String: "%14",
			Expected: true,
		},
		{
			String: "%15",
			Expected: true,
		},
		{
			String: "%16",
			Expected: true,
		},
		{
			String: "%17",
			Expected: true,
		},
		{
			String: "%18",
			Expected: true,
		},
		{
			String: "%19",
			Expected: true,
		},
		{
			String: "%1A",
			Expected: true,
		},
		{
			String: "%1B",
			Expected: true,
		},
		{
			String: "%1C",
			Expected: true,
		},
		{
			String: "%1D",
			Expected: true,
		},
		{
			String: "%1E",
			Expected: true,
		},
		{
			String: "%1F",
			Expected: true,
		},



		{
			String: "%1a",
			Expected: false,
		},
		{
			String: "%1b",
			Expected: false,
		},
		{
			String: "%1c",
			Expected: false,
		},
		{
			String: "%1d",
			Expected: false,
		},
		{
			String: "%1e",
			Expected: false,
		},
		{
			String: "%1f",
			Expected: false,
		},



		{
			String: "%20",
			Expected: true,
		},
		{
			String: "%21",
			Expected: true,
		},
		{
			String: "%22",
			Expected: true,
		},
		{
			String: "%23",
			Expected: true,
		},
		{
			String: "%24",
			Expected: true,
		},
		{
			String: "%25",
			Expected: true,
		},
		{
			String: "%26",
			Expected: true,
		},
		{
			String: "%27",
			Expected: true,
		},
		{
			String: "%28",
			Expected: true,
		},
		{
			String: "%29",
			Expected: true,
		},
		{
			String: "%2A",
			Expected: true,
		},
		{
			String: "%2B",
			Expected: true,
		},
		{
			String: "%2C",
			Expected: true,
		},
		{
			String: "%2D",
			Expected: true,
		},
		{
			String: "%2E",
			Expected: true,
		},
		{
			String: "%2F",
			Expected: true,
		},



		{
			String: "%2a",
			Expected: false,
		},
		{
			String: "%2b",
			Expected: false,
		},
		{
			String: "%2c",
			Expected: false,
		},
		{
			String: "%2d",
			Expected: false,
		},
		{
			String: "%2e",
			Expected: false,
		},
		{
			String: "%2f",
			Expected: false,
		},



		{
			String: "%30",
			Expected: true,
		},
		{
			String: "%31",
			Expected: true,
		},
		{
			String: "%32",
			Expected: true,
		},
		{
			String: "%33",
			Expected: true,
		},
		{
			String: "%34",
			Expected: true,
		},
		{
			String: "%35",
			Expected: true,
		},
		{
			String: "%36",
			Expected: true,
		},
		{
			String: "%37",
			Expected: true,
		},
		{
			String: "%38",
			Expected: true,
		},
		{
			String: "%39",
			Expected: true,
		},
		{
			String: "%3A",
			Expected: true,
		},
		{
			String: "%3B",
			Expected: true,
		},
		{
			String: "%3C",
			Expected: true,
		},
		{
			String: "%3D",
			Expected: true,
		},
		{
			String: "%3E",
			Expected: true,
		},
		{
			String: "%3F",
			Expected: true,
		},



		{
			String: "%3a",
			Expected: false,
		},
		{
			String: "%3b",
			Expected: false,
		},
		{
			String: "%3c",
			Expected: false,
		},
		{
			String: "%3d",
			Expected: false,
		},
		{
			String: "%3e",
			Expected: false,
		},
		{
			String: "%3f",
			Expected: false,
		},



		{
			String: "%40",
			Expected: true,
		},
		{
			String: "%41",
			Expected: true,
		},
		{
			String: "%42",
			Expected: true,
		},
		{
			String: "%43",
			Expected: true,
		},
		{
			String: "%44",
			Expected: true,
		},
		{
			String: "%45",
			Expected: true,
		},
		{
			String: "%46",
			Expected: true,
		},
		{
			String: "%47",
			Expected: true,
		},
		{
			String: "%48",
			Expected: true,
		},
		{
			String: "%49",
			Expected: true,
		},
		{
			String: "%4A",
			Expected: true,
		},
		{
			String: "%4B",
			Expected: true,
		},
		{
			String: "%4C",
			Expected: true,
		},
		{
			String: "%4D",
			Expected: true,
		},
		{
			String: "%4E",
			Expected: true,
		},
		{
			String: "%4F",
			Expected: true,
		},



		{
			String: "%4a",
			Expected: false,
		},
		{
			String: "%4b",
			Expected: false,
		},
		{
			String: "%4c",
			Expected: false,
		},
		{
			String: "%4d",
			Expected: false,
		},
		{
			String: "%4e",
			Expected: false,
		},
		{
			String: "%4f",
			Expected: false,
		},
	}

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	{
		var chars []byte = []byte{
			'.','-','_',
			'0','1','2','3','4','5','6','7','8','9',
			'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','w','x','y','z',
			'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z',
		}

		for i:=0; i<0xFF; i++ {
			var first byte = chars[randomness.Intn(len(chars))]

			var str string = string(first) + arbitrary.String()

			test := struct{
				String string
				Expected bool
			}{
				String: str,
				Expected: true,
			}

			tests = append(tests, test)
		}
	}

	{
		for i:=0; i<0xFF; i++ {
			var first string = arbitrary.PctEncoded()

			var str string = string(first) + arbitrary.String()

			test := struct{
				String string
				Expected bool
			}{
				String: str,
				Expected: true,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actual := syntax.HasPrefixIDChar(test.String)

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
