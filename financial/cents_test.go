package financial

import "testing"

func nilCentDict() CentDict {
	var n CentDict
	return n
}

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
		{"nil", nilCentDict(), false, "", 0},
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

func TestMergeWith(t *testing.T) {
	tests := []struct {
		name          string
		src, incoming CentDict
		expected      CentDict
	}{
		{
			name:     "both nil",
			src:      nilCentDict(),
			incoming: nilCentDict(),
			expected: CentDict{},
		},
		{
			name:     "incoming nil",
			incoming: nilCentDict(),
			src: CentDict{
				"a": 1000,
			},
			expected: CentDict{
				"a": 1000,
			},
		},
	}

	for _, test := range tests {
		test.src.MergeWith(test.incoming)

		if !test.src.Compare(test.expected) {
			t.Errorf("Testing %s.  Comparison failed. Expected %v; got %v", test.name, test.expected, test.src)
		}
	}
}

func TestRemoveTaxableSurcharge(t *testing.T) {
	tests := []struct {
		name          string
		netAmount     Cents
		surcharge     float64
		taxPercentage float64
		expected      Cents
	}{
		{
			name:          "all zero",
			netAmount:     0,
			surcharge:     0,
			taxPercentage: 0,
			expected:      0,
		},
		{
			name:          "zero surcharge",
			netAmount:     5000,
			surcharge:     0,
			taxPercentage: 0,
			expected:      5000,
		},
		{
			name:          "without tax",
			netAmount:     5500,
			surcharge:     10,
			taxPercentage: 0,
			expected:      5000,
		},

		{
			name:          "with tax",
			netAmount:     5600,
			surcharge:     10,
			taxPercentage: 20,
			expected:      5000,
		},
	}

	for _, test := range tests {
		res := test.netAmount.RemoveTaxableSurcharge(test.surcharge, test.taxPercentage)
		if res != test.expected {
			t.Errorf("Testing %s.  Comparison failed. Expected %v; got %v", test.name, test.expected, res)
		}
	}
}
