package main

import (
	"fmt"
	"golang-mono-micro/pkg/common/cmd"
	orders_app "golang-mono-micro/pkg/orders/application"
	orders_infra_orders "golang-mono-micro/pkg/orders/infrastructure/orders"
	orders_infra_payments "golang-mono-micro/pkg/orders/infrastructure/payments"
	orders_infra_product "golang-mono-micro/pkg/orders/infrastructure/shop"
	orders_private_http "golang-mono-micro/pkg/orders/interfaces/private/http"
	orders_public_http "golang-mono-micro/pkg/orders/interfaces/public/http"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Starting Orders Microservice");

	ctx := cmd.Context();
	r, closeFn := createOrderMicroservice();
	
	defer closeFn()

	server := &http.Server{
		Addr: os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR"),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed{
			panic(err)
		}
	}()

	<-ctx.Done();
	log.Println("Closing Orders Microservice");

	if err := server.Close(); err != nil{
		panic(err);
	}
}

func createOrderMicroservice()(router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"));

	shopHttpClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_PRODUCTS_SERVICE_ADDR"));
	ordersToPayQueue, err := orders_infra_payments.NewAMQPService(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
	)

	if err != nil {
		panic(err)
	}

	ordersRepo := orders_infra_orders.NewMemoryRepository();
	ordersService := orders_app.NewOrdersService(
		shopHttpClient,
		ordersToPayQueue,
		ordersRepo,
	)
	
	r := cmd.CreateRouter();

	orders_public_http.AddRoutes(r, ordersService, ordersRepo);
	orders_private_http.AddRoutes(r, ordersService, ordersRepo);

	return r, func() {
		err := ordersToPayQueue.Close();
		if err != nil {
			log.Printf("cannot close orders queue: %s", err)
		}
	}
}