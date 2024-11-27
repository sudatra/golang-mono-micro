package products

import "errors"

var ErrorNotFound = errors.New("product not found")

type Repository interface {
	Save(*Product) error
	ByID(ID) (*Product, error)
	AllProducts() ([]Product, error)
}

