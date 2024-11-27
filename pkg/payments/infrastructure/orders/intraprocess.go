package orders

import "golang-mono-micro/pkg/orders/interfaces/private/intraprocess"

type IntraprocessService struct {
	paymentsInterface intraprocess.OrdersInterface
}

func NewIntraprocessService(paymentsInterface intraprocess.OrdersInterface) IntraprocessService {
	return IntraprocessService{
		paymentsInterface: paymentsInterface,
	}
}

func (i IntraprocessService) MarkOrderAsPaid(orderID string) error {
	return i.paymentsInterface.MarkOrderAsPaid(orderID);
}