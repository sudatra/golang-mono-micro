package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"context"
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
	log.PrintLn("Closing Orders Microservice");

	if err := server.Close(); err != nil{
		panic(err);
	}
}

func createOrderMicroservice()(router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"));

	shopHttpClient := orders_infra_product.NewHttpClient(os.Getenv("SHOP_PRODUCTS_SERVICE_ADDR"));
	r := cmd.CreateRouter();

	orders_public_http.AddRoutes(r, ordersService, ordersRepo);
	orders_private_http.AddRoutes(r, ordersService, ordersRepo);

	return r, func() {}
}