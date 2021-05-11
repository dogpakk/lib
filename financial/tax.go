package financial

// Important
// This is an implementation of 'unit' method or 'early rounding' calculation
// as described in this post https://pakk.io/post/vat-rounding
// Make sure this is actually how you want to calculate tax before using

type TaxRoundingMethod uint

const (
	TaxRoundingMethodUnit TaxRoundingMethod = iota
	TaxRoundingMethodLine
	TaxRoundingMethodTotals
)

type TaxCalc struct {
	RoundingMethod TaxRoundingMethod
	LineQty        int
	TaxPercentage  float64

	UnitEx, LineEx, LineTax, LineInc Cents
}

type TaxCalcs []TaxCalc

func (tx *TaxCalc) blank() {
	tx.UnitEx = 0
	tx.LineEx = 0
	tx.LineTax = 0
	tx.LineInc = 0
}

// startFromEx basically resets the calculation, leaving the Ex as the starting point
// and setting inc to ex (i.e. a zero tax rate)
func (tx *TaxCalc) startFromUnitEx() {
	tx.LineTax = 0
	tx.LineEx = 0
	tx.LineInc = 0
}

func (tx *TaxCalc) startFromInc() {
	tx.UnitEx = 0
	tx.LineEx = 0
	tx.LineTax = 0
}

// AddTax adds tax for a line unit price and qty.  By default it uses the "line" rounding method,
// but can also use the "unit" method.  The "total" method is irrelevant here because
// line tax totals don't come into place in that case, so we just use the default line method as well.
func (tx *TaxCalc) AddTax() {
	if tx.RoundingMethod == TaxRoundingMethodUnit {
		tx.AddTaxUnitMethod()
		return
	}

	tx.AddTaxLineMethod()
}

// RemoveTax removes tax for a line unit price and qty.  By default it uses the "line" rounding method,
// but can also use the "unit" method.  The "total" method is irrelevant here because
// line tax totals don't come into place in that case, so we just use the default line method as well.
func (tx *TaxCalc) RemoveTax() {
	if tx.RoundingMethod == TaxRoundingMethodUnit {
		tx.RemoveTaxUnitMethod()
		return
	}

	tx.RemoveTaxLineMethod()
}

// AddTax goes from a unit ex price, qty and tax percentage to total line ex, tax and inc
func (tx *TaxCalc) AddTaxUnitMethod() {
	qty := tx.LineQty
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

	tx.LineEx = tx.UnitEx * Cents(qty)
	tx.LineTax = unitTax * Cents(qty)
	tx.LineInc = unitInc * Cents(qty)

	return
}

func (tx *TaxCalc) AddTaxLineMethod() {
	qty := tx.LineQty
	taxPercentage := tx.TaxPercentage

	// zero quantity means everything is zero
	if qty == 0 {
		tx.blank()
		return
	}

	// Reset
	tx.startFromUnitEx()

	// This is the line method, so multiply out the line ex total first
	// and use that as the basis of the tax calculation
	tx.LineEx = tx.UnitEx * Cents(qty)
	tx.LineTax = tx.LineEx.ByPercentage(taxPercentage)
	tx.LineInc = tx.LineEx + tx.LineTax

	return
}

func (tx *TaxCalc) RemoveTaxUnitMethod() {
	qty := tx.LineQty
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
	unitInc := tx.LineInc / Cents(qty)
	unitEx := unitInc.RemovePercentage(taxPercentage)
	unitTax := unitInc - unitEx

	tx.UnitEx = unitEx
	tx.LineEx = unitEx * Cents(qty)
	tx.LineTax = unitTax * Cents(qty)
}

func (tx *TaxCalc) RemoveTaxLineMethod() {
	qty := tx.LineQty
	taxPercentage := tx.TaxPercentage

	// zero quantity means everything is zero
	if qty == 0 {
		tx.blank()
		return
	}

	// Reset
	tx.startFromInc()

	tx.LineEx = tx.LineInc.RemovePercentage(taxPercentage)
	tx.LineTax = tx.LineInc - tx.LineEx
	tx.UnitEx = tx.LineEx / Cents(qty)
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
		lineEx := float64(tx.UnitEx) * float64(tx.LineQty)
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
