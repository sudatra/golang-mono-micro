package price

import "errors"

var (
	ErrorPriceTooLow = errors.New("price too low")
	ErrorInvalidCurrency = errors.New("invalid currency")
)

type Price struct {
	cents uint
	currency string
}

func NewPrice(cents uint, currency string) (Price, error) {
	if cents <= 0 {
		return Price{}, ErrorPriceTooLow
	}

	if len(currency) != 3 {
		return Price{}, ErrorInvalidCurrency
	}

	return Price{
		cents: cents,
		currency: currency,
	}, nil
}

func NewPricePanic(cents uint, currency string) Price {
	p, err := NewPrice(cents, currency);
	if err != nil {
		panic(err);
	}

	return p
}

func (p Price) Cents() uint {
	return p.cents
}

func (p Price) Currency() string {
	return p.currency
}