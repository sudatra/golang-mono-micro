package intraprocess

import (
	"golang-mono-micro/pkg/orders/application"
	"golang-mono-micro/pkg/orders/domain/orders"
)

type OrdersInterface struct {
	service application.OrdersService
}

func NewOrdersInterface(service application.OrdersService) OrdersInterface {
	return OrdersInterface{
		service: service,
	}
}

func (p OrdersInterface) MarkOrderAsPaid(orderID string) error {
	return p.service.MarkOrderAsPaid(application.MarkOrderAsPaidCommand{
		OrderID: orders.ID(orderID),
	})
}