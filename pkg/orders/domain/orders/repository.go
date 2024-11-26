package orders

import "errors"

var ErrorNotFound = errors.New("order not found");

type Repository interface {
	Save(*Order) error
	ByID(ID) (*Order, error)
}