package intraprocess

import (
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/payments/application"
	"log"
	"sync"
)

type OrderToProcess struct {
	ID string
	Price price.Price
}

type PaymentsInterface struct {
	orders <-chan OrderToProcess
	service application.PaymentsService
	orderProcessingWg *sync.WaitGroup
	runEnded chan struct{}
}

func NewPaymentsInterface(orders <-chan OrderToProcess, service application.PaymentsService) PaymentsInterface {
	return PaymentsInterface{
		orders: orders,
		service: service,
		orderProcessingWg: &sync.WaitGroup{},
		runEnded: make(chan struct{}, 1),
	}
}

func (p PaymentsInterface) Run() {
	defer func() {
		p.runEnded <- struct{}{}
	}()

	for order := range p.orders {
		p.orderProcessingWg.Add(1)
		
		go func(orderToPay OrderToProcess) {
			defer p.orderProcessingWg.Done()

			if err := p.service.InitializeOrderPayment(orderToPay.ID, orderToPay.Price); err != nil {
				log.Print("cannot initialize payment", err);
			}
		}(order)
	}
}

func (p PaymentsInterface) Close() {
	p.orderProcessingWg.Wait();
	<-p.runEnded
}