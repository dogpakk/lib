package str

import "testing"

func TestZeroPad(t *testing.T) {
	expected := "00000999"

	if res := ZeroPad("999", 8); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestZeroPad0(t *testing.T) {
	expected := "999"

	if res := ZeroPad("999", 0); res != expected {
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

func TestPrefixAndZeroPad0(t *testing.T) {
	expected := "PRE999"

	if res := PrefixAndZeroPad("999", "PRE", 0, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_TooShort(t *testing.T) {
	// With the specified count set to 5, and 'include' set to true,
	// there are not enough characters to fully express the prefix + ref combo
	// so default to the original
	expected := "PRE999"

	if res := PrefixAndZeroPad("999", "PRE", 5, true); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_NoPrefix(t *testing.T) {
	expected := "00000999"

	if res := PrefixAndZeroPad("999", "", 8, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_0NoPrefix(t *testing.T) {
	expected := "999"

	// shouldn't make any difference what we pass for 'include'
	if res := PrefixAndZeroPad("999", "", 0, true); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}

	if res := PrefixAndZeroPad("999", "", 0, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}
