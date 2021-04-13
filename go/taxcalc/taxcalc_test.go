package taxcalc_test

import (
	"math"
	"testing"

	"github.com/danblade/progressive_tax/go/taxcalc"
	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {

	for _, tt := range []struct {
		income float64
		tax    float64
	}{
		{
			income: 0,
			tax:    0,
		},
		{
			income: 10000,
			tax:    0,
		},
		{
			income: 10009,
			tax:    0,
		},
		{
			income: 10010,
			tax:    1,
		},
		{
			income: 12000,
			tax:    200,
		},
		{
			income: 56789,
			tax:    8697,
		},
		{
			income: 1234567,
			tax:    473326,
		},
	} {
		calculator := taxcalc.NewCalc([]taxcalc.Bracket{
			{
				Rate:  0,
				Lower: 0,
				Upper: 10000,
			},
			{
				Rate:  0.1,
				Lower: 10000,
				Upper: 30000,
			},
			{
				Rate:  0.25,
				Lower: 30000,
				Upper: 100000,
			},
			{
				Rate:  0.4,
				Lower: 100000,
				Upper: math.MaxFloat64,
			},
		})
		tax := calculator.Tax(tt.income)
		assert.Equal(t, tt.tax, tax)
	}
}
func TestTaxableIncome(t *testing.T) {
	for _, tt := range []struct {
		taxableIncome float64
		lower         float64
		upper         float64
		income        float64
	}{
		{0, 5, 10, 3},
		{0, 5, 10, 5},
		{3, 5, 10, 8},
		{5, 5, 10, 10},
		{5, 5, 10, 11},
	} {
		assert.EqualValues(t, tt.taxableIncome, taxcalc.TaxableIncome(tt.lower, tt.upper, tt.income))
	}

}
