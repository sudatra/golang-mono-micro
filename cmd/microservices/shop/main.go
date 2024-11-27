package main

import (
	"golang-mono-micro/pkg/common/cmd"
	"log"
	"net/http"
	"os"
	shop_infra_product "golang-mono-micro/pkg/shop/infrastructure/products"
	shop_interfaces_private_http "golang-mono-micro/pkg/shop/interfaces/private/http"
	shop_interfaces_public_http "golang-mono-micro/pkg/shop/interfaces/public/http"
	shop_app "golang-mono-micro/pkg/shop/application"
	shop "golang-mono-micro/pkg/shop"
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
	shopProductsService := shop_app.NewProductsService(shopProductRepo, shopProductRepo);

	if err := shop.LoadShopFixtures(shopProductsService); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter();

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo);
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo);

	return r; 
}
