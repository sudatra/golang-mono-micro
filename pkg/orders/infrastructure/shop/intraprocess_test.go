package shop_test

import (
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/orders/domain/orders"
	"golang-mono-micro/pkg/orders/infrastructure/shop"
	"golang-mono-micro/pkg/shop/interfaces/private/intraprocess"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOrderProductFromShopProduct(t *testing.T) {
	shopProduct := intraprocess.Product{
		ID: "123",
		Name: "name",
		Description: "desc",
		Price: price.NewPricePanic(42, "EUR"),
	}

	orderProduct, err := shop.OrderProductFromIntraprocess(shopProduct);
	assert.NoError(t, err)

	expectedOrderProduct, err := orders.NewProduct("123", "name", price.NewPricePanic(42, "EUR"));
	assert.NoError(t, err)

	assert.EqualValues(t, expectedOrderProduct, orderProduct)
}