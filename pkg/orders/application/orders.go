package application

import (
	"errors"
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/orders/domain/orders"
	"log"
)

type productsService interface {
	ProductsByID(id orders.ProductID) (orders.Product, error)
}

type paymentsService interface {
	InitializeOrderPayment(id orders.ID, price price.Price) error
}

type OrdersService struct {
	productsService productsService
	paymentsService paymentsService
	ordersRepository orders.Repository
}

type PlaceOrderCommand struct {
	OrderID orders.ID
	ProductID orders.ProductID
	Address PlaceOrderCommandAddress
}

type PlaceOrderCommandAddress struct {
	Name string
	Street string
	City string
	PostCode string
	Country string
}

type MarkOrderAsPaidCommand struct {
	OrderID orders.ID
}

func NewOrdersService(productsService productsService, paymentsService paymentsService, ordersRepository orders.Repository) OrdersService {
	return OrdersService{
		productsService: productsService,
		paymentsService: paymentsService,
		ordersRepository: ordersRepository,
	}
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {
	address, err := orders.NewAddress(
		cmd.Address.Name,
		cmd.Address.Street,
		cmd.Address.City,
		cmd.Address.PostCode,
		cmd.Address.Country,
	)

	if err != nil {
		return errors.New("invalid address");
	}

	product, err := s.productsService.ProductsByID(cmd.ProductID);
	if err != nil {
		return errors.New("cannot get the product");
	}

	newOrder, err := orders.NewOrder(cmd.OrderID, product, address);
	if err != nil {
		return errors.New("cannot place the order");
	}

	if err := s.ordersRepository.Save(newOrder); err != nil {
		return errors.New("cannot save order")
	}

	if err := s.paymentsService.InitializeOrderPayment(newOrder.ID(), newOrder.Product().Price()); err != nil {
		return errors.New("cannot initialize Payment")
	}

	log.Printf("order %s placed", cmd.OrderID);
	return nil;
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	o, err := s.ordersRepository.ByID(cmd.OrderID);
	if err != nil {
		return errors.New("cannot get order");
	}

	o.MarkAsPaid();
	if err := s.ordersRepository.Save(o); err != nil {
		return errors.New("cannot save order");
	}

	log.Printf("marked order as %s paid", cmd.OrderID);
	return nil;
}

func (s OrdersService) OrderById(id orders.ID) (orders.Order, error) {
	o, err := s.ordersRepository.ByID(id);
	if err != nil {
		return orders.Order{}, errors.New("cannot get order");
	}

	return *o, nil;
}