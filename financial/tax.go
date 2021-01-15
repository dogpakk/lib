package financial

// Important
// This is an implementation of 'unit' method or 'early rounding' calculation
// as described in this post https://pakk.io/post/vat-rounding
// Make sure this is actually how you want to calculate tax before using

type TaxCalc struct {
	Qty           int
	TaxPercentage float64

	UnitEx, Ex, Tax, Inc Cents
}

type TaxCalcs []TaxCalc

func (tx *TaxCalc) blank() {
	tx.UnitEx = 0
	tx.Ex = 0
	tx.Tax = 0
	tx.Inc = 0
}

// startFromEx basically resets the calculation, leaving the Ex as the starting point
// and setting inc to ex (i.e. a zero tax rate)
func (tx *TaxCalc) startFromUnitEx() {
	tx.Tax = 0
	tx.Ex = 0
	tx.Inc = 0
}

func (tx *TaxCalc) startFromInc() {
	tx.UnitEx = 0
	tx.Ex = 0
	tx.Tax = 0
}

// AddTax goes from a unit ex price, qty and tax percentage to total line ex, tax and inc
func (tx *TaxCalc) AddTax() {
	qty := tx.Qty
	taxPercentage := tx.TaxPercentage

	// zero quantity means everything is zero
	if qty == 0 {
		tx.blank()
		return
	}

	// Reset
	tx.startFromUnitEx()

	// Unit tax amount and inc tax are first calculated so they are correct
	// in and of themselves
	unitTax := tx.UnitEx.ByPercentage(taxPercentage)
	unitInc := tx.UnitEx + unitTax

	tx.Ex = tx.UnitEx * Cents(qty)
	tx.Tax = unitTax * Cents(qty)
	tx.Inc = unitInc * Cents(qty)

	return
}

func (tx *TaxCalc) RemoveTax() {
	qty := tx.Qty
	taxPercentage := tx.TaxPercentage

	// zero quantity means everything is zero
	if qty == 0 {
		tx.blank()
		return
	}

	// Reset
	tx.startFromInc()

	// Remeber, this is the unit method, so division by qty is first
	// which gets us to a unit inc
	unitInc := tx.Inc / Cents(qty)
	unitEx := unitInc.RemovePercentage(taxPercentage)
	unitTax := unitInc - unitEx

	tx.UnitEx = unitEx
	tx.Ex = unitEx * Cents(qty)
	tx.Tax = unitTax * Cents(qty)
}

func (txs TaxCalcs) CalcAggTaxRate() float64 {
	if len(txs) == 0 {
		return 0
	}

	// I don't use the normal VAT rounding technique here as it leads to roudning errors
	// being transmitted to the aggregate tax percetnage, so, e.g. multiple products with a 20%
	// VAT rate can end up having an aggregate of 20.02% or something like that.
	var totalEx, totalTax float64

	for _, tx := range txs {
		lineEx := float64(tx.UnitEx) * float64(tx.Qty)
		totalEx = totalEx + lineEx
		tax := lineEx * (tx.TaxPercentage / 100)
		totalTax = totalTax + tax
	}

	// Divide by zero protection
	if totalEx == 0 {
		return 0
	}

	return (totalTax / totalEx) * 100
}
