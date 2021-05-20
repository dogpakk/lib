package financial

import (
	"testing"

	"github.com/dogpakk/lib/slice"
)

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

func TestCentDictResetAllValues(t *testing.T) {
	tests := []struct {
		name string
		cd   CentDict
	}{
		{"blank", CentDict{}},
		{"one key already nil", CentDict{"one": 0}},
		{"two keys already nil", CentDict{"one": 0, "two": 0}},
		{"two keys already nil", CentDict{"one": 0, "two": 0}},
		{"one key", CentDict{"one": 99}},
		{"two keys mixed", CentDict{"one": 0, "two": 99}},
		{"two keys", CentDict{"one": 99, "two": 99}},
	}

	for _, test := range tests {
		test.cd.ResetAllValues()
		if test.cd.HasAnyNonZeroValues() {
			t.Errorf("Testing %s.  Expecting all values to be zero but got: %v", test.name, test.cd)
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

func TestKeyedCentDictAddToKey(t *testing.T) {
	tests := []struct {
		name     string
		existing KeyedCentDict
		k1, k2   string
		amount   Cents
		expected Cents
	}{
		{
			name:     "start with blank map, add nothing",
			existing: nilKeyedCentDict(),
			k1:       "K1",
			k2:       "K2",
			amount:   0,
			expected: 0,
		},
		{
			name:     "start with blank map, add something",
			existing: nilKeyedCentDict(),
			k1:       "K1",
			k2:       "K2",
			amount:   100,
			expected: 100,
		},
		{
			name: "start with top level keys as blank maps, add something",
			existing: KeyedCentDict{
				"K1":  CentDict{},
				"K1b": CentDict{},
			},
			k1:       "K1",
			k2:       "K2",
			amount:   100,
			expected: 100,
		},
		{
			name: "start with top level keys with existing, add something",
			existing: KeyedCentDict{
				"K1": CentDict{
					"K2": 100,
				},
				"K1b": CentDict{},
			},
			k1:       "K1",
			k2:       "K2",
			amount:   100,
			expected: 200,
		},
	}

	for _, test := range tests {
		test.existing.AddToKey(test.k1, test.k2, test.amount)
		if res, ok := test.existing[test.k1][test.k2]; !ok {
			t.Errorf("Testing %s.  Could not access map key", test.name)
		} else if res != test.expected {
			t.Errorf("Testing %s.  Incorrect result. Expected %v; got %v", test.name, test.expected, res)
		}
	}
}

func TestKeyedCentDictAllSecondLevelKeys(t *testing.T) {
	tests := []struct {
		name     string
		src      KeyedCentDict
		expected []string
	}{
		{
			name:     "blank",
			src:      KeyedCentDict{},
			expected: []string{},
		},
		{
			name: "blank second levels",
			src: KeyedCentDict{
				"K1":  CentDict{},
				"K1b": CentDict{},
			},
			expected: []string{},
		},
		{
			name: "2 top level, 1 second level",
			src: KeyedCentDict{
				"K1": CentDict{
					"K2": 100,
				},
				"K1b": CentDict{},
			},
			expected: []string{"K2"},
		},
		{
			name: "2 top level, 2 second level",
			src: KeyedCentDict{
				"K1": CentDict{
					"K2": 100,
					"K3": 100,
				},
				"K1b": CentDict{},
			},
			expected: []string{"K2", "K3"},
		},
		{
			name: "2 top level, 2 second level, all duplicates",
			src: KeyedCentDict{
				"K1": CentDict{
					"K2": 100,
					"K3": 100,
				},
				"K1b": CentDict{
					"K2": 100,
					"K3": 100,
				},
			},
			expected: []string{"K2", "K3"},
		},
		{
			name: "2 top level, 2 second level, 1 duplicate",
			src: KeyedCentDict{
				"K1": CentDict{
					"K2": 100,
					"K3": 100,
				},
				"K1b": CentDict{
					"K2":  100,
					"K3a": 100,
				},
			},
			expected: []string{"K2", "K3", "K3a"},
		},
	}

	for _, test := range tests {
		res := test.src.AllSecondLevelKeys()
		if len(res) != len(test.expected) {
			t.Errorf("Testing %s.  Incorrect result. Expected %v; got %v", test.name, len(test.expected), len(res))
		}

		if !slice.CompareStringSlicesOrderIrrelevant(res, test.expected) {
			t.Errorf("Testing %s.  Result slices don't match. Expected %v; got %v", test.name, test.expected, res)
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

func TestFormatAndParsePrices(t *testing.T) {
	tests := []struct {
		c        Cents
		expected string
	}{
		{0, "0.00"},
		{1, "0.01"},
		{10, "0.10"},
		{100, "1.00"},
		{1000, "10.00"},
		{10000, "100.00"},
		{0, "0.00"},
		{-1, "-0.01"},
		{-10, "-0.10"},
		{-100, "-1.00"},
		{-1000, "-10.00"},
		{-10000, "-100.00"},
		{1949, "19.49"}, // Tests rounding - 19.49 parses to float as 19.4899999999
	}

	for _, test := range tests {
		res := test.c.FormatAsPrice()
		if res != test.expected {
			t.Errorf("Formatting. Comparison failed. Expected %v; got %v", test.expected, res)
		}

		parsed, err := ParseCentsFromPriceString(test.expected)
		if err != nil {
			t.Errorf("Parsing failed: %s", err)
		}

		if parsed != test.c {
			t.Errorf("Parsing. Comparison failed. Expected %v; got %v", test.c, parsed)
		}
	}
}
