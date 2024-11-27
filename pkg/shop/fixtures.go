package shop

import (
	"errors"
	shop_app "golang-mono-micro/pkg/shop/application"
)

func LoadShopFixtures(productsService shop_app.ProductsService) error {
	err := productsService.AddProduct(shop_app.AddProductCommand{
		ID: "1",
		Name: "Product 1",
		Description: "Random Product 1",
		PriceCents: 500,
		Currency: "INR",
	})

	if err != nil {
		return errors.New("unable to create shop fixture");
	}

	return productsService.AddProduct(shop_app.AddProductCommand{
		ID: "2",
		Name: "Product 2",
		Description: "Random Product 2",
		PriceCents: 580,
		Currency: "GBP",
	})
}