package orders

import (
	"errors"
	"golang-mono-micro/pkg/common/price"
)

type ProductID string
var ErrorEmptyProductID = errors.New("empty product id");

type Product struct {
	id ProductID
	name string
	price price.Price
}

func NewProduct(id ProductID, name string, price price.Price) (Product, error) {
	if len(id) == 0 {
		return Product{}, ErrorEmptyProductID
	}

	return Product{
		id: id,
		name: name,
		price: price,
	}, nil
}

func (p Product) ID() ProductID {
	return p.id;
}

func (p Product) Name() string {
	return p.name;
}

func (p Product) Price() price.Price {
	return p.price;
}