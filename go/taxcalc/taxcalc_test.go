package taxcalc_test

import (
	"fmt"
	"testing"

	"github.com/danblade/progressive_tax/go/taxcalc"
	"github.com/shopspring/decimal"
	// "github.com/danblade/progressive_tax/go/taxcalc"
	// "github.com/shopspring/decimal"
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
			tax:    decimal.RequireFromString("0"),
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
			tax:    decimal.RequireFromString("8697"),
		},
		{
			income: decimal.RequireFromString("1234567"),
			tax:    decimal.RequireFromString("473326"),
		},
	} {
		// 	taxcalc.NewCalc([]taxcalc.Bracket{
		// 		{
		// 			Rate:  decimal.RequireFromString("0.1"),
		// 			Lower: decimal.RequireFromString("0"),
		// 			Upper: decimal.RequireFromString("10000"),
		// 		},
		// 		{
		// 			Rate:  decimal.RequireFromString("0.175"),
		// 			Lower: decimal.RequireFromString("10000"),
		// 			Upper: decimal.RequireFromString("30000"),
		// 		},
		// 		{
		// 			Rate:  decimal.RequireFromString("0.25"),
		// 			Lower: decimal.RequireFromString("30000"),
		// 			Upper: decimal.RequireFromString("100000"),
		// 		},
		// 		{
		// 			Rate:  decimal.RequireFromString("0.4"),
		// 			Lower: decimal.RequireFromString("100000"),
		// 		},
		// 	}, tt.income)
		fmt.Printf("Tax: $%v Source Income: %v\n", taxcalc.Tax(tt.income), tt.income)
	}
}
