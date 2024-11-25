package products

import (
	"errors"
	"golang-mono-micro/pkg/common/price"
)

type ID string

var (
	ErrorEmptyID = errors.New("empty product id")
	ErrorEmptyName = errors.New("empty product name")
)

type Product struct {
	id ID
	name string
	description string
	price price.Price
}

func NewProduct(id ID, name string, description string, price price.Price) (*Product, error) {
	if len(id) == 0 {
		return nil, ErrorEmptyID
	}

	if len(name) == 0 {
		return nil, ErrorEmptyName
	}

	return &Product{
		id,
		name,
		description,
		price,
	}, nil
}

func (p Product) ID() ID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Description() string {
	return p.description
}

func (p Product) Price() price.Price {
	return p.price
}