package str

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewCleanString(t *testing.T) {
	testCases := []struct {
		name         string
		in, expected string
	}{
		{"blank string", "", ""},
		{"simple lower case", "test", "test"},
		{"simple upper case", "TEST", "test"},
		{"lower case with space to left", " test", "test"},
		{"lower case with space to right", "test ", "test"},
		{"lower case with space on both sides", " test ", "test"},
		{"upper case with space to left", " TEST", "test"},
		{"upper case with space to right", "TEST ", "test"},
		{"upper case with space on both sides", " TEST ", "test"},
		{"upper case with space on both sides and inside (respect internal space)", " TE ST ", "te st"},
	}

	for _, test := range testCases {
		out := NewCleanString(test.in)
		if string(out) != test.expected {
			t.Errorf("Testing %s failed.  Expected %s; got %s", test.name, test.expected, string(out))
		}

		// A little String output test for when a CleanString is forcefully created with a bad string
		badCS := CleanString(test.in)
		if badCS.String() != test.expected {
			t.Errorf("Forced CS: Testing %s failed.  Expected %s; got %s", test.name, test.expected, badCS.String())
		}

		// JSON Unmarshalling
		var newCS CleanString
		if err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, test.in)), &newCS); err != nil {
			t.Errorf("Testing %s, unmarshalling, failed with error: %s.", test.name, err)
		}

		if string(newCS) != test.expected {
			t.Errorf("Testing %s, unmarhsalling failed.  Expected %s; got %s", test.name, test.expected, string(newCS))
		}
	}

}
