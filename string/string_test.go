package string

import "testing"

func TestZeroPad(t *testing.T) {
	expected := "00000999"

	if res := ZeroPad("999", 8); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_ExcludePrefixFromCount(t *testing.T) {
	expected := "PRE00000999"

	if res := PrefixAndZeroPad("999", "PRE", 8, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_IncludePrefixInCount(t *testing.T) {
	expected := "PRE00999"

	if res := PrefixAndZeroPad("999", "PRE", 8, true); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}
