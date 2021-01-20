package financial

import "testing"

func TestCentRatio(t *testing.T) {
	tests := []struct {
		name     string
		a, b     Cents
		expected float64
	}{
		{"a and b = 0", 0, 0, 0},
		{"a = 0", 0, 1, 0},
		{"b = 0", 1, 0, 0},
		{"divide by 1", 2, 1, 2},
		{"simple exact divsion", 1, 2, 0.5},
		{"simple exact divsion", 1, 2, 0.5},
		{"recursive", 2, 3, 0.6666666666666666},
	}

	for _, test := range tests {
		res := CentRatio(test.a, test.b)
		if res != test.expected {
			t.Errorf("Error testing %s. Expected %v, got %v", test.name, test.expected, res)
		}
	}
}

func TestHasOneKey(t *testing.T) {
	tests := []struct {
		name           string
		src            CentDict
		expectedBool   bool
		expectedString string
		expectedCents  Cents
	}{
		{"blank", CentDict{}, false, "", 0},
		{"one key", CentDict{"myKey": 99}, true, "myKey", 99},
		{"two keys", CentDict{"myKey": 99, "anotherKey": 999}, false, "", 0},
	}

	for _, test := range tests {
		b, s, c := test.src.HasOneKey()

		if b != test.expectedBool {
			t.Errorf("Testing %s.  Incorrect bool. Expected %v; got %v", test.name, test.expectedBool, b)
		}

		if s != test.expectedString {
			t.Errorf("Testing %s.  Incorrect bool. Expected %v; got %v", test.name, test.expectedString, s)
		}

		if c != test.expectedCents {
			t.Errorf("Testing %s.  Incorrect bool. Expected %v; got %v", test.name, test.expectedCents, c)
		}
	}

}
