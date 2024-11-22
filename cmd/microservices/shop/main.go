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
	log.Println("Starting Shop Microservice");

	ctx := cmd.Context();
	r := createShopMicroservice();

	server := &http.Server{
		Addr: os.Getenv("SHOP_PRODUCTS_SERVICE_BIND_ADDR"),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed{
			panic(err)
		}
	}()

	<-ctx.Done();
	log.Println("Closing Shop Microservice");

	if err := server.Close(); err != nil{
		panic(err)
	}
}

func createShopMicroservice()(router *chi.Mux) {
	shopProductRepo := shop_infra_product.NewMemoryRepository();
	r := cmd.CreateRouter();

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo);
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo);

	return r; 
}
