package shop

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/orders/domain/orders"
	http_interface "golang-mono-micro/pkg/shop/interfaces/private/http"
	"io"
	"net/http"
)

type HTTPClient struct {
	address string
}

func NewHTTPClient(address string) HTTPClient {
	return HTTPClient{
		address: address,
	}
}

func OrderProductPriceFromHTTP(priceView http_interface.PriceView) (price.Price, error) {
	return price.NewPrice(priceView.Cents, priceView.Currency);
}

func OrderProductFromHTTP(shop_product http_interface.ProductView) (orders.Product, error) {
	productPrice, err := OrderProductPriceFromHTTP(shop_product.Price);
	if err != nil {
		return orders.Product{}, errors.New("cannot decode price");
	}

	return orders.NewProduct(
		orders.ProductID(shop_product.ID),
		shop_product.Name,
		productPrice,
	)
}

func (h HTTPClient) ProductByID(id orders.ProductID) (orders.Product, error) {
	res, err := http.Get(fmt.Sprintf("%s/product/%s", h.address, id));
	if err != nil {
		return orders.Product{}, errors.New("get shop failed")
	}

	defer func() {
		_ = res.Body.Close()
	}();

	b, err := io.ReadAll(res.Body);
	if err != nil {
		return orders.Product{}, errors.New("cannot read response")
	}

	productView := http_interface.ProductView{};
	if err := json.Unmarshal(b, &productView); err != nil {
		return orders.Product{}, errors.New("cannot decode response");
	}

	return OrderProductFromHTTP(productView);
}