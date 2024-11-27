package orders_test

import (
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/orders/domain/orders"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createOrderContent(t *testing.T) (orders.Product, orders.Address) {
	productPrice, err := price.NewPrice(10, "USD");
	assert.NoError(t, err);

	orderProduct, err := orders.NewProduct("1", "foo", productPrice);
	assert.NoError(t, err);

	orderAddress, err := orders.NewAddress("test Name", "test Street", "test City", "test PostCode", "test Country");
	assert.NoError(t, err);

	return orderProduct, orderAddress
}

func NewOrderTest(t *testing.T) {
	orderProduct, orderAddress := createOrderContent(t);
	testOrder, err := orders.NewOrder("1", orderProduct, orderAddress);

	assert.NoError(t, err);
	assert.EqualValues(t, orderProduct, testOrder.Product());
	assert.EqualValues(t, orderAddress, testOrder.Address());
	assert.False(t, testOrder.Paid());
}

func TestNewOrder_empty_id(t *testing.T) {
	orderProduct, orderAddress := createOrderContent(t);
	_, err := orders.NewOrder("", orderProduct, orderAddress);

	assert.EqualValues(t, orders.ErrorEmptyOrderID, err);
}

func TestNewOrder_MarkAsPaid(t *testing.T) {
	orderProduct, orderAddress := createOrderContent(t);
	testOrder, err := orders.NewOrder("1", orderProduct, orderAddress);

	assert.NoError(t, err);
	assert.False(t, testOrder.Paid());
	testOrder.MarkAsPaid();
	assert.True(t, testOrder.Paid());
}