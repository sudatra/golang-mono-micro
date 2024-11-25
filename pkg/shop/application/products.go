package application

import (
	"errors"
	"golang-mono-micro/pkg/common/price"
	products "golang-mono-micro/pkg/shop/domain"
)

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

type AddProductCommand struct {
	ID          string
	Name        string
	Description string
	PriceCents  uint
	Currency    string
}

func NewProductsService(repo products.Repository, readModel productReadModel) ProductsService {
	return ProductsService{
		repo: repo,
		readModel: readModel,
	}
}

func (s ProductsService) AllProducts() ([]products.Product, error) {
	return s.readModel.AllProducts()
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
	price, err := price.NewPrice(cmd.PriceCents, cmd.Currency);
	if err != nil {
		return errors.New("invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)
	if err != nil {
		return errors.New("cannot create product")
	}

	if err := s.repo.Save(p); err != nil {
		return errors.New("cannot save product")
	}
	
	return nil;
}