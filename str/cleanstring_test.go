package str

import "testing"

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
		if out.String() != test.expected {
			t.Errorf("Testing %s failed.  Expected %s; got %s", test.name, test.expected, out)
		}
	}

}
