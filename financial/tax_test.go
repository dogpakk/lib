package financial

import "testing"

func TestAddTax(t *testing.T) {
	tests := []struct {
		name     string
		in       TaxCalc
		expected TaxCalc
	}{
		// BASIC EDGE CASE TESTS

		// LineQty 1
		{
			name: "qty 1, blank and zero tax pc",
			in: TaxCalc{
				LineQty: 1,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, blank and positive tax pc",
			in: TaxCalc{
				LineQty:       1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, no ex, results are already filled with junk",
			in: TaxCalc{
				LineQty: 1,
				LineTax: 99,
				LineInc: 88,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, simple calculation",
			in: TaxCalc{
				UnitEx:        100,
				LineQty:       1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  100,
				LineTax: 25,
				LineInc: 125,
			},
		},
		{
			name: "qty 1, simple calculation - prefilled junk",
			in: TaxCalc{
				LineQty:       1,
				TaxPercentage: 25,
				UnitEx:        100,
				LineTax:       99,
				LineInc:       999,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  100,
				LineTax: 25,
				LineInc: 125,
			},
		},
		{
			name: "qty 1, zero tax pc but positive ex",
			in: TaxCalc{
				LineQty:       1,
				TaxPercentage: 0,
				UnitEx:        100,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineTax: 0,
				LineInc: 100,
				LineEx:  100,
			},
		},
		// LineQty 0 - always expect a nil result
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
				LineTax: 99,
				LineInc: 88,
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
				LineTax:       99,
				LineInc:       999,
			},
			expected: TaxCalc{},
		},
		// Mutliple LineQty
		{
			name: "qty 3, simple calculation",
			in: TaxCalc{
				TaxPercentage: 25,
				UnitEx:        100,
				LineQty:       3,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  300,
				LineTax: 75,
				LineInc: 375,
			},
		},
		{
			name: "qty 3, zero tax pc but positive ex",
			in: TaxCalc{
				LineQty:       3,
				TaxPercentage: 0,
				UnitEx:        100,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineTax: 0,
				LineInc: 300,
				LineEx:  300,
			},
		},

		// MORE ADVANCED CALCULATION TESTS INVOLVING ROUNDING AND MULTIPLES
		{
			name: "unambiguous round downwards",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.25,
				LineQty:       1,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  100,
				LineTax: 17,
				LineInc: 117,
			},
		},
		{
			name: "unambiguous round upwards",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.75,
				LineQty:       1,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  100,
				LineTax: 18,
				LineInc: 118,
			},
		},
		{
			name: "exact 0.5 should round up",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.5,
				LineQty:       1,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  100,
				LineTax: 18,
				LineInc: 118,
			},
		},
		{
			// Because of the particular choice of numbers here, the unit method and line method produce
			// calcs with a 1 cent difference.  This test proves that we are using the unit method
			name: "mismatch between line and unit - unit method",
			in: TaxCalc{
				UnitEx:        100,
				TaxPercentage: 17.4,
				LineQty:       3,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  300,
				LineTax: 51,
				LineInc: 351,
			},
		},
		{
			// Because of the particular choice of numbers here, the unit method and line method produce
			// calcs with a 1 cent difference.  This test proves that we are using the unit method
			name: "mismatch between line and unit - line method",
			in: TaxCalc{
				RoundingMethod: TaxRoundingMethodLine,
				UnitEx:         100,
				TaxPercentage:  17.4,
				LineQty:        3,
			},
			expected: TaxCalc{
				RoundingMethod: TaxRoundingMethodLine,
				UnitEx:         100,
				LineEx:         300,
				LineTax:        52,
				LineInc:        352,
			},
		},
	}

	for _, test := range tests {
		test.in.AddTax()

		if test.in.UnitEx != test.expected.UnitEx ||
			test.in.LineEx != test.expected.LineEx ||
			test.in.LineTax != test.expected.LineTax ||
			test.in.LineInc != test.expected.LineInc {
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

		// LineQty 1
		{
			name: "qty 1, blank and zero tax pc",
			in: TaxCalc{
				LineQty: 1,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, blank and positive tax pc",
			in: TaxCalc{
				LineQty:       1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, no inc, results are already filled with junk",
			in: TaxCalc{
				LineQty: 1,
				LineTax: 99,
				LineEx:  88,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 1, simple calculation",
			in: TaxCalc{
				LineInc:       100,
				LineQty:       1,
				TaxPercentage: 25,
			},
			expected: TaxCalc{
				UnitEx:  80,
				LineEx:  80,
				LineTax: 20,
				LineInc: 100,
			},
		},
		{
			name: "qty 1, simple calculation - prefilled junk",
			in: TaxCalc{
				LineQty:       1,
				TaxPercentage: 25,
				LineInc:       100,
				LineEx:        99,
				UnitEx:        999,
			},
			expected: TaxCalc{
				UnitEx:  80,
				LineEx:  80,
				LineTax: 20,
				LineInc: 100,
			},
		},
		{
			name: "qty 1, zero tax pc but positive ex",
			in: TaxCalc{
				LineQty:       1,
				TaxPercentage: 0,
				LineInc:       100,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineTax: 0,
				LineInc: 100,
				LineEx:  100,
			},
		},
		// LineQty 0 - always expect null result
		{
			name: "qty 0, blank and zero tax pc",
			in: TaxCalc{
				LineQty: 0,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, blank and positive tax pc",
			in: TaxCalc{
				LineQty:       0,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, no inc, results are already filled with junk",
			in: TaxCalc{
				LineQty: 0,
				LineTax: 99,
				LineEx:  88,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, simple calculation",
			in: TaxCalc{
				LineInc:       100,
				LineQty:       0,
				TaxPercentage: 25,
			},
			expected: TaxCalc{},
		},
		{
			name: "qty 0, simple calculation - prefilled junk",
			in: TaxCalc{
				LineQty:       0,
				TaxPercentage: 25,
				LineInc:       100,
				LineEx:        99,
				UnitEx:        999,
			},
			expected: TaxCalc{},
		},
		// Mutliple LineQty
		{
			name: "qty 3, simple calculation",
			in: TaxCalc{
				LineQty:       3,
				LineInc:       375,
				TaxPercentage: 25,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  300,
				LineTax: 75,
				LineInc: 375,
			},
		},
		{
			name: "qty 3, zero tax pc but positive ex",
			in: TaxCalc{
				LineQty:       3,
				TaxPercentage: 0,
				LineInc:       300,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineTax: 0,
				LineInc: 300,
				LineEx:  300,
			},
		},

		// MORE ADVANCED CALCULATION TESTS INVOLVING ROUNDING AND MULTIPLES
		{
			name: "unambiguous round updwards",
			in: TaxCalc{
				LineInc:       10000,
				TaxPercentage: 17.25,
				LineQty:       1,
			},
			expected: TaxCalc{
				UnitEx:  8529,
				LineEx:  8529,
				LineTax: 1471,
				LineInc: 10000,
			},
		},
		{
			name: "Rounding with multiples",
			in: TaxCalc{
				LineInc:       30300,
				TaxPercentage: 17.4,
				LineQty:       3,
			},
			expected: TaxCalc{
				UnitEx:  8603,
				LineEx:  25809,
				LineTax: 4491,
				LineInc: 30300,
			},
		},
		{
			// Because of the particular choice of numbers here, the unit method and line method produce
			// calcs with a 1 cent difference.  This test proves that we are using the unit method
			name: "mismatch between line and unit - unit method",
			in: TaxCalc{
				LineInc:       351,
				TaxPercentage: 17.4,
				LineQty:       3,
			},
			expected: TaxCalc{
				UnitEx:  100,
				LineEx:  300,
				LineTax: 51,
				LineInc: 351,
			},
		},
		{
			// Because of the particular choice of numbers here, the unit method and line method produce
			// calcs with a 1 cent difference.  This test proves that we are using the unit method
			name: "mismatch between line and unit - unit method",
			in: TaxCalc{
				RoundingMethod: TaxRoundingMethodLine,
				LineInc:        352,
				TaxPercentage:  17.4,
				LineQty:        3,
			},
			expected: TaxCalc{
				RoundingMethod: TaxRoundingMethodLine,
				UnitEx:         100,
				LineEx:         300,
				LineTax:        52,
				LineInc:        352,
			},
		},
	}

	for _, test := range tests {
		test.in.RemoveTax()

		if test.in.UnitEx != test.expected.UnitEx ||
			test.in.LineEx != test.expected.LineEx ||
			test.in.LineTax != test.expected.LineTax ||
			test.in.LineInc != test.expected.LineInc {
			t.Fatalf("Testing '%s' - got mismatch. Expected %v; got %v",
				test.name,
				test.expected,
				test.in,
			)
		}
	}
}
