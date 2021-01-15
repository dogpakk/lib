package financial

import "testing"

func TestAddTax(t *testing.T) {
	tests := []struct {
		name     string
		in       TaxCalc
		expected TaxCalc
	}{
		// BASIC EDGE CASE TESTS

		// Qty 1
		{
			name: "qty 1, blank and zero tax pc",
			in: TaxCalc{
				Qty: 1,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, blank and positive tax pc",
			in: TaxCalc{
				Qty:           1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, no ex, results are already filled with junk",
			in: TaxCalc{
				Qty: 1,
				Tax: 99,
				Inc: 88,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, simple calculation",
			in: TaxCalc{
				UnitEx:        100,
				Qty:           1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     100,
				Tax:    25,
				Inc:    125,
			},
		},
		{
			name: "qty 1, simple calculation - prefilled junk",
			in: TaxCalc{
				Qty:           1,
				TaxPercentage: 25,
				UnitEx:        100,
				Tax:           99,
				Inc:           999,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     100,
				Tax:    25,
				Inc:    125,
			},
		},
		{
			name: "qty 1, zero tax pc but positive ex",
			in: TaxCalc{
				Qty:           1,
				TaxPercentage: 0,
				UnitEx:        100,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Tax:    0,
				Inc:    100,
				Ex:     100,
			},
		},
		// Qty 0 - always expect a nil result
		{
			name:     "qty 0, blank and zero tax pc",
			in:       TaxCalc{},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, blank and positive tax pc",
			in: TaxCalc{
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, no ex, results are already filled with junk",
			in: TaxCalc{
				Tax: 99,
				Inc: 88,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, simple calculation",
			in: TaxCalc{
				TaxPercentage: 25,
				UnitEx:        100,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, simple calculation - prefilled junk",
			in: TaxCalc{
				TaxPercentage: 25,
				UnitEx:        100,
				Tax:           99,
				Inc:           999,
			},
			expected: TaxCalc{},
		},
		// Mutliple Qty
		{
			name: "qty 3, simple calculation",
			in: TaxCalc{
				TaxPercentage: 25,
				UnitEx:        100,
				Qty:           3,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     300,
				Tax:    75,
				Inc:    375,
			},
		},
		{
			name: "qty 3, zero tax pc but positive ex",
			in: TaxCalc{
				Qty:           3,
				TaxPercentage: 0,
				UnitEx:        100,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Tax:    0,
				Inc:    300,
				Ex:     300,
			},
		},

		// MORE ADVANCED CALCULATION TESTS INVOLVING ROUNDING AND MULTIPLES
		{
			name: "unambiguous round downwards",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.25,
				Qty:           1,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     100,
				Tax:    17,
				Inc:    117,
			},
		},
		{
			name: "unambiguous round upwards",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.75,
				Qty:           1,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     100,
				Tax:    18,
				Inc:    118,
			},
		},
		{
			name: "exact 0.5 should round up",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.5,
				Qty:           1,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     100,
				Tax:    18,
				Inc:    118,
			},
		},
		{
			// Because of the particular choice of numbers here, the unit method and line method produce
			// calcs with a 1 cent difference.  This test proves that we are using the unit method
			name: "mismatch between line and unit",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.4,
				Qty:           3,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     300,
				Tax:    51,
				Inc:    351,
			},
		},
	}

	for _, test := range tests {
		test.in.AddTax()

		if test.in.UnitEx != test.expected.UnitEx ||
			test.in.Ex != test.expected.Ex ||
			test.in.Tax != test.expected.Tax ||
			test.in.Inc != test.expected.Inc {
			t.Fatalf("Testing '%s' - got mismatch. Expected %v; got %v",
				test.name,
				test.expected,
				test.in,
			)
		}
	}
}

func TestRemoveTax(t *testing.T) {
	tests := []struct {
		name     string
		in       TaxCalc
		expected TaxCalc
	}{
		// BASIC EDGE CASE TESTS

		// Qty 1
		{
			name: "qty 1, blank and zero tax pc",
			in: TaxCalc{
				Qty: 1,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, blank and positive tax pc",
			in: TaxCalc{
				Qty:           1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, no inc, results are already filled with junk",
			in: TaxCalc{
				Qty: 1,
				Tax: 99,
				Ex:  88,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, simple calculation",
			in: TaxCalc{
				Inc:           100,
				Qty:           1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{
				UnitEx: 80,
				Ex:     80,
				Tax:    20,
				Inc:    100,
			},
		},
		{
			name: "qty 1, simple calculation - prefilled junk",
			in: TaxCalc{
				Qty:           1,
				TaxPercentage: 25,
				Inc:           100,
				Ex:            99,
				UnitEx:        999,
			},
			expected: TaxCalc{
				UnitEx: 80,
				Ex:     80,
				Tax:    20,
				Inc:    100,
			},
		},
		{
			name: "qty 1, zero tax pc but positive ex",
			in: TaxCalc{
				Qty:           1,
				TaxPercentage: 0,
				Inc:           100,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Tax:    0,
				Inc:    100,
				Ex:     100,
			},
		},
		// Qty 0 - always expect null result
		{
			name: "qty 0, blank and zero tax pc",
			in: TaxCalc{
				Qty: 0,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, blank and positive tax pc",
			in: TaxCalc{
				Qty:           0,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, no inc, results are already filled with junk",
			in: TaxCalc{
				Qty: 0,
				Tax: 99,
				Ex:  88,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, simple calculation",
			in: TaxCalc{
				Inc:           100,
				Qty:           0,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, simple calculation - prefilled junk",
			in: TaxCalc{
				Qty:           0,
				TaxPercentage: 25,
				Inc:           100,
				Ex:            99,
				UnitEx:        999,
			},
			expected: TaxCalc{},
		},
		// Mutliple Qty
		{
			name: "qty 3, simple calculation",
			in: TaxCalc{
				Qty:           3,
				Inc:           375,
				TaxPercentage: 25,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Ex:     300,
				Tax:    75,
				Inc:    375,
			},
		},
		{
			name: "qty 3, zero tax pc but positive ex",
			in: TaxCalc{
				Qty:           3,
				TaxPercentage: 0,
				Inc:           300,
			},
			expected: TaxCalc{
				UnitEx: 100,
				Tax:    0,
				Inc:    300,
				Ex:     300,
			},
		},

		// MORE ADVANCED CALCULATION TESTS INVOLVING ROUNDING AND MULTIPLES
		{
			name: "unambiguous round updwards",
			in: TaxCalc{
				Inc:           10000,
				TaxPercentage: 17.25,
				Qty:           1,
			},
			expected: TaxCalc{
				UnitEx: 8529,
				Ex:     8529,
				Tax:    1471,
				Inc:    10000,
			},
		},
		{
			name: "Rounding with multiples",
			in: TaxCalc{
				Inc:           30300,
				TaxPercentage: 17.4,
				Qty:           3,
			},
			expected: TaxCalc{
				UnitEx: 8603,
				Ex:     25809,
				Tax:    4491,
				Inc:    30300,
			},
		},
	}

	for _, test := range tests {
		test.in.RemoveTax()

		if test.in.UnitEx != test.expected.UnitEx ||
			test.in.Ex != test.expected.Ex ||
			test.in.Tax != test.expected.Tax ||
			test.in.Inc != test.expected.Inc {
			t.Fatalf("Testing '%s' - got mismatch. Expected %v; got %v",
				test.name,
				test.expected,
				test.in,
			)
		}
	}
}
