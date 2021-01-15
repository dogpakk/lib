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
