package price_test

import (
	"golang-mono-micro/pkg/common/price"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrice(t *testing.T) {
	testCases := []struct {
		Name string
		Cents uint
		Currency string
		ExpectedError error
	} {
		{
			Name: "valid",
			Cents: 10,
			Currency: "EUR",
		},
		{
			Name: "invalid_cents",
			Cents: 0,
			Currency: "EUR",
			ExpectedError: price.ErrorPriceTooLow,
		},
		{
			Name: "empty_currency",
			Cents: 10,
			Currency: "",
			ExpectedError: price.ErrorInvalidCurrency,
		},
		{
			Name: "invalid_currency_length",
			Cents: 10,
			Currency: "US",
			ExpectedError: price.ErrorInvalidCurrency,
		},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			_, err := price.NewPrice(c.Cents, c.Currency);
			assert.EqualValues(t, c.ExpectedError, err);
		})
	}
}