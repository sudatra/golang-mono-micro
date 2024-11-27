package orders_test

import (
	"golang-mono-micro/pkg/orders/domain/orders"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	testCases := []struct {
		TestName string
		Name string
		Street string
		City string
		PostCode string
		Country string
		ExpectedError bool
	} {
		{
			TestName:    "valid",
			Name:        "test Name",
			Street:      "test Street",
			City:        "test City",
			PostCode:    "test PostCode",
			Country:     "test Country",
			ExpectedError: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			address, err := orders.NewAddress(c.Name, c.Street, c.City, c.PostCode, c.Country)

			if c.ExpectedError {
				assert.Error(t, err);
			} else {
				assert.NoError(t, err);

				assert.EqualValues(t, c.Name, address.Name());
				assert.EqualValues(t, c.Street, address.Street());
				assert.EqualValues(t, c.City, address.City());
				assert.EqualValues(t, c.PostCode, address.PostCode());
				assert.EqualValues(t, c.Country, address.Country());
			}
		})
	}
}