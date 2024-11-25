package http

import (
	"golang-mono-micro/pkg/common/price"
	products "golang-mono-micro/pkg/shop/domain"
	"net/http"
	common_http "golang-mono-micro/pkg/common/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
}

type productsResource struct {
	readModel productsReadModel
}

type ProductView struct {
	ID string `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Price PriceView `json:"price"`
}

type PriceView struct {
	Cents uint `json:"cents"`
	Currency string `json:"currency"`
}

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

func PriceViewFromPrice(p price.Price) PriceView {
	return PriceView{
		Cents: p.Cents(),
		Currency: p.Currency(),
	}
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AllProducts();
	if err != nil {
		_ = render.Render(w, r, common_http.ErrorInternal(err))
		return;
	}

	view := []ProductView{};
	for _, product := range products {
		view = append(view, ProductView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			PriceViewFromPrice(product.Price()),
		})
	}

	render.Respond(w, r, view);
}