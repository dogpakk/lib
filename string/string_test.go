package string

import "testing"

func TestZeroPad(t *testing.T) {
	expected := "00000999"

	if res := ZeroPad("999", 8); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}
