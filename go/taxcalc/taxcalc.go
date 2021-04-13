package taxcalc

import (
	"math"
)

type Bracket struct {
	Rate  float64
	Lower float64
	Upper float64
}

type Calc struct {
	brackets []Bracket
}

func NewCalc(bracket []Bracket) *Calc {
	return &Calc{brackets: bracket}
}

func (c *Calc) Tax(income float64) float64 {
	var tax float64
	for _, bracket := range c.brackets {
		taxableIncome := TaxableIncome(bracket.Lower, bracket.Upper, income)
		if taxableIncome > 0 {
			tax += taxableIncome * bracket.Rate
		}
	}
	return math.Floor(tax)
}

func TaxableIncome(lower, upper, income float64) float64 {
	if income < lower {
		return 0
	} else if income > upper {
		return upper - lower
	}
	return income - lower
}
