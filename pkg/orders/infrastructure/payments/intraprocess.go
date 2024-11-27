package payments

import (
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/orders/domain/orders"
	"golang-mono-micro/pkg/payments/interfaces/intraprocess"
)

type IntraprocessService struct {
	orders chan<- intraprocess.OrderToProcess
}

func NewIntraprocessService(ordersChannel chan<- intraprocess.OrderToProcess) IntraprocessService {
	return IntraprocessService{
		orders: ordersChannel,
	}
}

func (i IntraprocessService) InitializeOrderPayment(id orders.ID, price price.Price) error {
	i.orders <- intraprocess.OrderToProcess{
		ID: string(id),
		Price: price,
	}

	return nil
}