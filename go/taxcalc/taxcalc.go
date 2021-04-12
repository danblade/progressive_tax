package taxcalc

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Bracket struct {
	Rate  decimal.Decimal
	Lower decimal.Decimal
	Upper decimal.Decimal
}

type Calc struct {
	brackets []Bracket
}

func NewCalc(bracket []Bracket) *Calc {
	return &Calc{brackets: bracket}
}

func (c *Calc) Tax(income decimal.Decimal) decimal.Decimal {
	first := c.brackets[0].Upper.Sub(c.brackets[0].Lower).Mul(c.brackets[0].Rate)
	second := c.brackets[1].Upper.Sub(c.brackets[1].Lower).Mul(c.brackets[1].Rate)
	third := c.brackets[2].Upper.Sub(c.brackets[2].Lower).Mul(c.brackets[2].Rate)

	fmt.Printf("\n\n1st section tax: %v\n2nd section tax: %v\n3rd section tax: %v\n", first, second, third)
	if income.GreaterThan(c.brackets[3].Lower) {
		fmt.Printf("\n%v\n", income.Sub(c.brackets[3].Lower).Mul(c.brackets[3].Rate))
		fmt.Println(income.Sub(c.brackets[3].Lower).Mul(c.brackets[3].Rate).Add(first).Add(second).Add(third))
	}
	var nextLower decimal.Decimal
	_ = nextLower
	var tax decimal.Decimal
	for i := range c.brackets {
		// fmt.Println(tax)
		var taxableIncome decimal.Decimal
		if i+1 < len(c.brackets) {
			nextLower = c.brackets[i+1].Lower
		}
		if income.GreaterThan(c.brackets[i].Lower) {
			if income.GreaterThan(nextLower) {
				taxableIncome = income.Sub(nextLower).Sub(c.brackets[i].Lower)

			} else {
				taxableIncome = income.Sub(c.brackets[i].Lower)
				// if i+1 < len(c.brackets) {
				// 	if income.GreaterThan(c.brackets[i+1].Lower) {
				// 		taxableIncome = c.brackets[i+1].Lower
				// 	}
			}
			tax = taxableIncome.Mul(c.brackets[i].Rate).Add(tax)
		}
		fmt.Printf("\nTax after %v\n\n", tax)
	}
	fmt.Printf("\nTax: %v, income: %v\n", tax, income)
	return tax
}
