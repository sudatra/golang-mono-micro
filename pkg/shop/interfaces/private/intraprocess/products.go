package intraprocess

import (
	"errors"
	"golang-mono-micro/pkg/common/price"
	products "golang-mono-micro/pkg/shop/domain"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       price.Price
}

type ProductInterface struct {
	repo products.Repository
}

func NewProductInterface(repo products.Repository) ProductInterface {
	return ProductInterface{
		repo: repo,
	}
}

func ProductFromDomainProduct(domainProduct products.Product) Product {
	return Product{
		string(domainProduct.ID()),
		domainProduct.Name(),
		domainProduct.Description(),
		domainProduct.Price(),
	}
}

func (i ProductInterface) ProductByID(id string) (Product, error) {
	domainProduct, err := i.repo.ByID(products.ID(id));
	if err != nil {
		return Product{}, errors.New("cannot get product");
	}

	return ProductFromDomainProduct(*domainProduct), nil;
}