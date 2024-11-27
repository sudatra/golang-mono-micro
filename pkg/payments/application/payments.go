package application

import (
	"golang-mono-micro/pkg/common/price"
	"log"
	"time"
)

type ordersService interface {
	MarkOrderAsPaid(orderID string) error
}

type PaymentsService struct {
	ordersService ordersService
}

func NewPaymentsService(ordersService ordersService) PaymentsService {
	return PaymentsService{
		ordersService: ordersService,
	}
}

func (p PaymentsService) PostOrderPayment(orderID string) error {
	log.Printf("payment for order %s done, marking order as paid", orderID);
	return p.ordersService.MarkOrderAsPaid(orderID);
}

func (p PaymentsService) InitializeOrderPayment(orderID string, price price.Price) error {
	log.Printf("Initializing payment for order %s", orderID);

	go func() {
		time.Sleep(time.Millisecond * 500);
		if err := p.PostOrderPayment(orderID); err != nil {
			log.Printf("cannot post order payment: %s", err)
		}
	}()

	return nil;
}