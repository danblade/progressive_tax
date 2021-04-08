package taxcalc

import (
	"github.com/shopspring/decimal"
)

func Tax(income decimal.Decimal) decimal.Decimal {
	if income.GreaterThanOrEqual(decimal.RequireFromString("100000")) {

		return decimal.RequireFromString(
			"0.1").Mul(
			decimal.RequireFromString("20000")).Add(
			decimal.RequireFromString("0.25").Mul(
				decimal.RequireFromString("70000"))).Add(
			decimal.RequireFromString("0.4").Mul(
				income.Sub(decimal.RequireFromString("100000")),
			),
		)
	} else if income.GreaterThanOrEqual(decimal.RequireFromString("30000")) {

		return decimal.RequireFromString(
			"0.1").Mul(
			decimal.RequireFromString("20000")).Add(
			decimal.RequireFromString("0.25").Mul(
				income.Sub(decimal.RequireFromString("30000")),
			),
		)
	} else if income.GreaterThanOrEqual(decimal.RequireFromString("10000")) {
		return decimal.RequireFromString("0.1").Mul(
			income.Sub(decimal.RequireFromString("10000")),
		)
	}
	return decimal.Zero
}

// Works from : https://www.reddit.com/r/dailyprogrammer/comments/cdieag/20190715_challenge_379_easy_progressive_taxation/ew1dn5n?utm_source=share&utm_medium=web2x&context=3
// func Tax(income int) float64 {
// 	if income > 100000 {
// 		return (0.1 * float64(20000)) + (0.25 * float64(70000)) + (0.4 * math.Abs(float64(100000-income)))
// 	} else if income > 30000 {
// 		return (0.1 * float64(20000)) + (0.25 * math.Abs(float64(30000-income)))
// 	} else if income > 10000 {
// 		return 0.1 * math.Abs(float64(10000-income))
// 	} else {
// 		return 0
// 	}
// }

// type Bracket struct {
// 	Rate  decimal.Decimal
// 	Lower decimal.Decimal
// 	Upper decimal.Decimal
// }

// type Calc struct {
// 	tax    decimal.Decimal
// 	income decimal.Decimal
// }

// func NewCalc(bracket []Bracket, income decimal.Decimal) *Calc {
// 	// income := decimal.RequireFromString("12345")

// 	for i := range bracket {
// 		fmr.Println(Calc.DoTheMath(bracket))

// 	}
// 	return &Calc{}
// }

// func (c *Calc) DoTheMath(income decimal.Decimal) decimal.Decimal {
// 	var payment decimal.Decimal
// 	if income.GreaterThan(bracket[i].Lower) && income.LessThanOrEqual(bracket[i].Upper) {
// 		fmt.Printf("\ninfo: %v\n\n", lastLineTax)
// 		fmt.Println(bracket[i].Rate)
// 		fmt.Println(bracket[i].Lower, " < ", income, " and <= ", bracket[i].Upper)
// 		fmt.Println(bracket[i].Rate.Mul(income.Sub(bracket[i].Lower)).Round(2), " amount of tax for this bracket.")
// 		totalTax = bracket[i].Rate.Mul(income.Sub(bracket[i].Lower)).Round(2)
// 		lastLineTax = totalTax.Add(lastLineTax)
// 		fmt.Printf("\nTotal Tax: %v\n\n", lastLineTax)
// 	}

// 	return payment
// }
