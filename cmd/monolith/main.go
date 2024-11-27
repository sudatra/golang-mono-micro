package main

import (
	"golang-mono-micro/pkg/common/cmd"
	orders_app "golang-mono-micro/pkg/orders/application"
	orders_infra_orders "golang-mono-micro/pkg/orders/infrastructure/orders"
	orders_infra_payments "golang-mono-micro/pkg/orders/infrastructure/payments"
	orders_infra_product "golang-mono-micro/pkg/orders/infrastructure/shop"
	orders_interfaces_intraprocess "golang-mono-micro/pkg/orders/interfaces/private/intraprocess"
	orders_interfaces_http "golang-mono-micro/pkg/orders/interfaces/public/http"
	payments_app "golang-mono-micro/pkg/payments/application"
	payments_infra_orders "golang-mono-micro/pkg/payments/infrastructure/orders"
	payments_interfaces_intraprocess "golang-mono-micro/pkg/payments/interfaces/intraprocess"
	"golang-mono-micro/pkg/shop"
	shop_app "golang-mono-micro/pkg/shop/application"
	shop_infra_product "golang-mono-micro/pkg/shop/infrastructure/products"
	shop_interfaces_intraprocess "golang-mono-micro/pkg/shop/interfaces/private/intraprocess"
	shop_interfaces_http "golang-mono-micro/pkg/shop/interfaces/public/http"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Monolith Started");

	ctx := cmd.Context();
	ordersToPay := make(chan payments_interfaces_intraprocess.OrderToProcess);
	router, paymentsInterface := createMonolith(ordersToPay);

	go paymentsInterface.Run();

	server := &http.Server{
		Addr: os.Getenv("SHOP_MONOLITH_BIND_ADDR"),
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Printf("Monolith is listening on %s", server.Addr);

	<-ctx.Done()
	log.Printf("Closing Monolith");

	if err := server.Close(); err != nil {
		panic(err)
	}

	close(ordersToPay);
	paymentsInterface.Close();
}

func createMonolith(ordersToPay chan payments_interfaces_intraprocess.OrderToProcess) (*chi.Mux, payments_interfaces_intraprocess.PaymentsInterface) {
	shopProductsRepo := shop_infra_product.NewMemoryRepository();
	shopProductsService := shop_app.NewProductsService(shopProductsRepo, shopProductsRepo);
	shopProductIntraprocessInterface := shop_interfaces_intraprocess.NewProductInterface(shopProductsRepo);

	ordersRepo := orders_infra_orders.NewMemoryRepository();
	ordersService := orders_app.NewOrdersService(
		orders_infra_product.NewIntraprocessService(shopProductIntraprocessInterface),
		orders_infra_payments.NewIntraprocessService(ordersToPay),
		ordersRepo,
	)
	ordersIntraprocessInterface := orders_interfaces_intraprocess.NewOrdersInterface(ordersService)

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewIntraprocessService(ordersIntraprocessInterface),
	)
	paymentsIntraprocessInterface := payments_interfaces_intraprocess.NewPaymentsInterface(ordersToPay, paymentsService);

	if err := shop.LoadShopFixtures(shopProductsService); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter();

	shop_interfaces_http.AddRoutes(r, shopProductsRepo);
	orders_interfaces_http.AddRoutes(r, ordersService, ordersRepo);

	return r, paymentsIntraprocessInterface
}