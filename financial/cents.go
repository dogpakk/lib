package financial

import (
	"math"
	"math/rand"
	"strconv"
	"strings"

	pmath "github.com/dogpakk/lib/math"
	"github.com/dogpakk/lib/slice"
)

type Cents int

func percentageToDecimal(pc float64) float64 {
	return pc / 100
}

func percentageToDecimalIncrementer(pc float64) float64 {
	return 1 + percentageToDecimal(pc)
}

func (c Cents) LimitTo(n Cents) Cents {
	if c > n {
		return n
	}

	return c
}

func (c Cents) Min(n Cents) Cents {
	if c < n {
		return n
	}

	return c
}

func (c Cents) ByQty(qty int) Cents {
	return c * Cents(qty)
}

func (c Cents) DivideByQty(qty int) Cents {
	return Cents(math.Round(float64(c) / float64(qty)))
}

func (c Cents) ByFloat(multiplier float64) Cents {
	return Cents(math.Round(float64(c) * multiplier))
}

func (c Cents) ByPercentage(pc float64) Cents {
	return c.ByFloat(percentageToDecimal(pc))
}

func (c Cents) RemovePercentage(pc float64) Cents {
	return c.ByFloat(1 / percentageToDecimalIncrementer(pc))
}

func (c Cents) CalcPercentageDiscount(pc float64) Cents {
	return c.ByPercentage(pc)
}

// RemoveTaxableSurcharge works backwards from a net total that might include a surcharge which might itself be taxable
func (c Cents) RemoveTaxableSurcharge(surchargePercentage, taxPercentage float64) Cents {
	if surchargePercentage == 0 {
		return c
	}

	return c.RemovePercentage(surchargePercentage * percentageToDecimalIncrementer(taxPercentage))

}

func (c Cents) AddTax(taxPercentage float64) TaxCalc {
	// Rounding method makes no difference for single unit
	return c.AddTaxMultipleQuantity(1, taxPercentage, TaxRoundingMethodUnit)
}

func (c Cents) AddTaxMultipleQuantity(qty int, taxPercentage float64, roundingMethod TaxRoundingMethod) TaxCalc {
	taxCalc := TaxCalc{
		RoundingMethod: roundingMethod,
		LineQty:        qty,
		UnitEx:         c,
		TaxPercentage:  taxPercentage,
	}

	taxCalc.AddTax()
	return taxCalc
}

func (c Cents) RemoveTax(taxPercentage float64) TaxCalc {
	// Rounding method makes no difference for single unit
	return c.RemoveTaxMultipleQuantity(1, taxPercentage, TaxRoundingMethodUnit)
}

func (c Cents) RemoveTaxMultipleQuantity(qty int, taxPercentage float64, roundingMethod TaxRoundingMethod) TaxCalc {
	taxCalc := TaxCalc{
		RoundingMethod: roundingMethod,
		LineQty:        qty,
		LineInc:        c,
		TaxPercentage:  taxPercentage,
	}

	taxCalc.RemoveTax()
	return taxCalc
}

func (c Cents) SplitHundredths() (Cents, Cents) {
	remainder := c % 100
	return remainder, c - remainder
}

func (c Cents) RoundToNearestPretty(target Cents) Cents {
	// if the target is greater than 100,
	// we'll take the modulus by 100
	if target > 100 {
		target = target % 100
	}

	if c <= target {
		return target
	}

	if c <= 100 {
		return 100 + target
	}

	_, base := c.SplitHundredths()

	return base + target
}

func (c Cents) FormatAsPrice() (res string) {
	// If the number is negative, we'll convert it to positive for processing
	// and readd back in the sign right at the end
	isNegative := c < 0
	if isNegative {
		c = -c
	}

	s := strconv.Itoa(int(c))

	l := len(s)

	switch l {
	case 0:
		res = "0.00"
	case 1:
		res = "0.0" + s
	case 2:
		res = "0." + s
	default:
		prefix := s[:l-2]
		suffix := s[l-2:]
		res = strings.Join([]string{prefix, suffix}, ".")
	}

	if isNegative {
		res = "-" + res
	}

	return
}

func RandCents(n int) Cents {
	return Cents(rand.Intn(n))
}

func FromFloatPrice(f float64) Cents {
	return Cents(math.Round(100 * f))
}

func ParseCentsFromPriceString(s string) (Cents, error) {
	// We need this to deal with either "." or "," as unit separator
	// but we will NOT allow thousands separators here (as they are a human thing, not a machine thing)
	s = strings.Replace(s, ",", ".", 1)

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return FromFloatPrice(f), nil
}

func MustParseCentsFromPriceString(s string) Cents {
	cents, _ := ParseCentsFromPriceString(s)
	return cents
}

func CentRatio(a, b Cents) float64 {
	return pmath.IntRatio(int(a), int(b))
}

func CentRatioPercentage(a, b Cents) float64 {
	return pmath.IntRatioPercentage(int(a), int(b))
}

// CentDict is a string keyed map -> Cents, which is a data structure I find myself using a lot
// A good example is a map of currency balances, e.g.:
// GBP -> 456777
// EUR -> 322944
type CentDict map[string]Cents

func CompareCentDicts(cd1, cd2 CentDict) bool {
	for k1, v1 := range cd1 {
		if v1 > 0 {
			v2, ok := cd2[k1]
			if !ok {
				return false
			}

			if v1 != v2 {
				return false
			}
		}
	}

	return true
}

// IsZero will be true for a CentDict that is either nil,
// or has ANY values that are not 0
func (cd CentDict) IsZero() bool {
	if cd == nil {
		return true
	}

	for _, v := range cd {
		if v != 0 {
			return false
		}
	}

	return true
}

// Average will calculate the average of EACH KEY using a single divisor
func (cd CentDict) Average(n int) CentDict {
	if n == 0 || n == 1 {
		return cd
	}

	res := CentDict{}
	for k, v := range cd {
		if v > 0 {
			res[k] = v.DivideByQty(n)
		}
	}

	return res
}

func (cd CentDict) ByQty(n int) CentDict {
	if n == 0 || n == 1 {
		return cd
	}

	res := CentDict{}
	for k, v := range cd {
		res[k] = v.ByQty(n)
	}

	return res
}

func (cd CentDict) HasPositiveVals() bool {
	for _, v := range cd {
		if v > 0 {
			return true
		}
	}

	return false
}

func (cd CentDict) HasNegativeVals() bool {
	for _, v := range cd {
		if v < 0 {
			return true
		}
	}

	return false
}

func (cd CentDict) HasOneKey() (bool, string, Cents) {
	if len(cd) > 1 {
		return false, "", 0
	}

	for k, v := range cd {
		return true, k, v
	}

	return false, "", 0
}

func (cd CentDict) Compare(cd1 CentDict) bool {
	// test the map backwards and forwards to make sure they are really the same
	return CompareCentDicts(cd, cd1) && CompareCentDicts(cd1, cd)
}

func (cd CentDict) Invert() {
	// If map is nil, we can just return
	if cd == nil {
		return
	}

	for k, amount := range cd {
		cd[k] = -amount
	}
}

// AddToKey and MergeWith cannot be called on nil maps
// I can't find a way to deal neatly with this problem
// as trying to initialise the map in the function call
// causes the initial reference to be lost so the caller
// doesn't see the results of the function.
// Perhaps there's a way to initalise the map in place using
// some reflect code, but for now, just be sure
// NOT TO CALL THESE ON NIL MAPS

func (cd CentDict) AddToKey(k string, amount Cents) {
	// TODO: nil

	if existing, ok := cd[k]; ok {
		cd[k] = existing + amount
	} else {
		cd[k] = amount
	}
}

func (cd CentDict) MergeWith(incoming CentDict) {
	// If the incoming is nil, we can just return
	if incoming == nil {
		return
	}

	for k, amount := range incoming {
		if existing, ok := cd[k]; ok {
			cd[k] = existing + amount
		} else {
			cd[k] = amount
		}
	}
}

// KeyedCentDict is a map of maps: a string keyed map -> CentDict
// Useful, for example, to build up a report of regular CentDicts keyed by some variable
// E.g. A sales report by Item Sku
type KeyedCentDict map[string]CentDict

func nilKeyedCentDict() KeyedCentDict {
	return KeyedCentDict{}
}

func (kcd KeyedCentDict) AddToKey(k1, k2 string, amount Cents) {
	if existing, ok := kcd[k1]; ok {
		existing.AddToKey(k2, amount)
		kcd[k1] = existing
	} else {
		kcd[k1] = CentDict{k2: amount}
	}
}

func (kcd KeyedCentDict) AllSecondLevelKeys() (res []string) {
	for _, v1 := range kcd {
		for k := range v1 {
			res = slice.StringSliceAddToSet(k, res)
		}
	}

	return
}

func (kcd KeyedCentDict) AllTopLevelKeys() (res []string) {
	for k := range kcd {
		res = slice.StringSliceAddToSet(k, res)
	}

	return
}

// KeyedCentDict2 is a map of map of maps: a string keyed mao of KeyedCentDict
// An example use case is in grouping report sections

type KeyedCentDict2 map[string]KeyedCentDict

func nilKeyedCentDict2() KeyedCentDict2 {
	return KeyedCentDict2{}
}

func (kcd KeyedCentDict2) AddToKey(k1, k2, k3 string, amount Cents) {
	if existing, ok := kcd[k1]; ok {
		existing.AddToKey(k2, k3, amount)
		kcd[k1] = existing
	} else {
		kcd[k1] = KeyedCentDict{
			k2: CentDict{k3: amount},
		}
	}
}

func (kcd KeyedCentDict2) AllThirdLevelKeys() (res []string) {
	for _, v1 := range kcd {
		for _, v2 := range v1 {
			for k := range v2 {
				res = slice.StringSliceAddToSet(k, res)
			}
		}
	}

	return
}

func (kcd KeyedCentDict2) AllSecondLevelKeys() (res []string) {
	for _, v1 := range kcd {
		for k := range v1 {
			res = slice.StringSliceAddToSet(k, res)
		}
	}

	return
}

func (kcd KeyedCentDict2) AllTopLevelKeys() (res []string) {
	for k := range kcd {
		res = slice.StringSliceAddToSet(k, res)
	}

	return
}
