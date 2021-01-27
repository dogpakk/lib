package financial

import (
	"math"
	"math/rand"
	"strconv"
	"strings"

	pmath "github.com/dogpakk/lib/math"
)

type Cents int

func (c Cents) LimitTo(n Cents) Cents {
	if c > n {
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
	return c.ByFloat(pc / 100)
}

func (c Cents) RemovePercentage(pc float64) Cents {
	return c.ByFloat(1 / (1 + (pc / 100)))
}

func (c Cents) CalcPercentageDiscount(pc float64) Cents {
	return c.ByPercentage(pc)
}

func (c Cents) AddTax(taxPercentage float64) (Cents, Cents) {
	return c.AddTaxMultipleQuantity(1, taxPercentage)
}

func (c Cents) AddTaxMultipleQuantity(qty int, taxPercentage float64) (Cents, Cents) {
	taxCalc := TaxCalc{
		Qty:           qty,
		UnitEx:        c,
		TaxPercentage: taxPercentage,
	}

	taxCalc.AddTax()
	return taxCalc.Tax, taxCalc.Inc
}

func (c Cents) RemoveTax(taxPercentage float64) (Cents, Cents) {
	return c.RemoveTaxMultipleQuantity(1, taxPercentage)
}

func (c Cents) RemoveTaxMultipleQuantity(qty int, taxPercentage float64) (Cents, Cents) {
	taxCalc := TaxCalc{
		Qty:           qty,
		Inc:           c,
		TaxPercentage: taxPercentage,
	}

	taxCalc.RemoveTax()
	return taxCalc.Tax, taxCalc.Ex
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

	return
}

func RandCents(n int) Cents {
	return Cents(rand.Intn(n))
}

func ParseCentsFromPriceString(s string) (Cents, error) {
	// We need this to deal with either "." or "," as unit separator
	// but we will NOT allow thousands separators here (as they are a human thing, not a machine thing)
	s = strings.Replace(s, ",", ".", 1)

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return Cents(100 * f), nil
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

func (cd CentDict) HasPositiveKeys() bool {
	for k, v := range cd {
		if v > 0 {
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

func (cd CentDict) AddToKey(k string, amount Cents) {
	if existing, ok := cd[k]; ok {
		cd[k] = existing + amount
	} else {
		cd[k] = amount
	}
}

func (cd CentDict) MergeWith(incoming CentDict) {
	for k, amount := range incoming {
		if existing, ok := cd[k]; ok {
			cd[k] = existing + amount
		} else {
			cd[k] = amount
		}
	}
}

func (cd CentDict) Invert() {
	for k, amount := range cd {
		cd[k] = -amount
	}
}

// CentDict2 is a map of maps: a string keyed map -> CentDict
type CentDict2 map[string]CentDict
