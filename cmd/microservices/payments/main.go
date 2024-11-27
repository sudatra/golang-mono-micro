package main

import (
	"fmt"
	"golang-mono-micro/pkg/common/cmd"
	payments_app "golang-mono-micro/pkg/payments/application"
	payments_infra_orders "golang-mono-micro/pkg/payments/infrastructure/orders"
	"golang-mono-micro/pkg/payments/interfaces/amqp"
	"log"
	"os"
)

func main() {
	log.Println("Starting Payments Microservice");
	defer log.Println("Closing Payments Microservice");

	ctx := cmd.Context();
	paymentsInterface := createPaymentsMicroservice();

	if err := paymentsInterface.Run(ctx); err != nil{
		panic(err)
	}
}

func createPaymentsMicroservice() amqp.PaymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"));

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)
	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
	)

	if err != nil{
		panic(err)
	}

	return paymentsInterface
}