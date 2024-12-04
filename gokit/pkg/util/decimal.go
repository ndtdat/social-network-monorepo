package util

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func NullableDecimalToString(decimal *decimal.Decimal) string {
	if decimal == nil {
		return ""
	}

	return decimal.String()
}

func MustParseNullableDecimal(amount string) *decimal.Decimal {
	if amount == "" {
		return nil
	}

	decimal, err := decimal.NewFromString(amount)
	if err != nil {
		panic(fmt.Sprintf("cannot parse decimal for %s", amount))
	}

	return &decimal
}

func MustParseDecimal(amount string) decimal.Decimal {
	decimal, err := decimal.NewFromString(amount)
	if err != nil {
		panic(fmt.Sprintf("cannot parse decimal for %s", amount))
	}

	return decimal
}

func MustParseDecimalWithDefault(amount string) decimal.Decimal {
	if amount == "" {
		return decimal.Zero
	}

	decimal, err := decimal.NewFromString(amount)
	if err != nil {
		panic(fmt.Sprintf("cannot parse decimal for %s", amount))
	}

	return decimal
}
