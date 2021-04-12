package taxcalc_test

import (
	"testing"

	"github.com/danblade/progressive_tax/go/taxcalc"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {

	for _, tt := range []struct {
		income decimal.Decimal
		tax    decimal.Decimal
	}{
		{
			income: decimal.RequireFromString("0"),
			tax:    decimal.RequireFromString("0"),
		},
		{
			income: decimal.RequireFromString("10000"),
			tax:    decimal.RequireFromString("0"),
		},
		{
			income: decimal.RequireFromString("10009"),
			tax:    decimal.RequireFromString("0.9"),
		},
		{
			income: decimal.RequireFromString("10010"),
			tax:    decimal.RequireFromString("1"),
		},
		{
			income: decimal.RequireFromString("12000"),
			tax:    decimal.RequireFromString("200"),
		},
		{
			income: decimal.RequireFromString("56789"),
			tax:    decimal.RequireFromString("8697.25"),
		},
		{
			income: decimal.RequireFromString("1234567"),
			tax:    decimal.RequireFromString("473326.8"),
		},
	} {
		calculator := taxcalc.NewCalc([]taxcalc.Bracket{
			{
				Rate:  decimal.RequireFromString("0"),
				Lower: decimal.RequireFromString("0"),
				Upper: decimal.RequireFromString("10000"),
			},
			{
				Rate:  decimal.RequireFromString("0.1"),
				Lower: decimal.RequireFromString("10000"),
				Upper: decimal.RequireFromString("30000"),
			},
			{
				Rate:  decimal.RequireFromString("0.25"),
				Lower: decimal.RequireFromString("30000"),
				Upper: decimal.RequireFromString("100000"),
			},
			{
				Rate:  decimal.RequireFromString("0.4"),
				Lower: decimal.RequireFromString("100000"),
			},
		})
		tax := calculator.Tax(tt.income)
		assert.Truef(t, tax.Equal(tt.tax), "%v %v", tax, tt.tax)
	}
}
