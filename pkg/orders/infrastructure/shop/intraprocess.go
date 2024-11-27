package shop

import (
	"errors"
	"golang-mono-micro/pkg/orders/domain/orders"
	"golang-mono-micro/pkg/shop/interfaces/private/intraprocess"
)

type IntraProcessService struct {
	intraprocessInterface intraprocess.ProductInterface
}

func NewIntraprocessService(intraprocessInterface intraprocess.ProductInterface) IntraProcessService {
	return IntraProcessService{
		intraprocessInterface: intraprocessInterface,
	}
}

func OrderProductFromIntraprocess(shopProduct intraprocess.Product) (orders.Product, error) {
	return orders.NewProduct(
		orders.ProductID(shopProduct.ID),
		shopProduct.Name,
		shopProduct.Price,
	)
}

func (i IntraProcessService) ProductByID(id orders.ProductID) (orders.Product, error) {
	shopProduct, err := i.intraprocessInterface.ProductByID(string(id));
	if err != nil {
		return orders.Product{}, errors.New("unable to fetch product by id")
	}

	return OrderProductFromIntraprocess(shopProduct)
}