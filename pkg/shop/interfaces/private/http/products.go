package http

import (
	common_http "golang-mono-micro/pkg/common/http"
	"golang-mono-micro/pkg/common/price"
	products_domain "golang-mono-micro/pkg/shop/domain"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type productsResource struct {
	repo products_domain.Repository
}

type PriceView struct {
	Cents uint `json:"cents"`
	Currency string `json:"currency"`
}

type ProductView struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Price PriceView `json:"price"`
}

func AddRoutes(router *chi.Mux, repo products_domain.Repository) {
	resource := productsResource{repo: repo};
	router.Get("/products/{id}", resource.Get)
}

func PriceViewFromPrice(p price.Price) PriceView {
	return PriceView{
		Cents: p.Cents(),
		Currency: p.Currency(),
	}
} 

func (p productsResource) Get(w http.ResponseWriter, r *http.Request) {
	product, err := p.repo.ByID(products_domain.ID(chi.URLParam(r, "id")));
	if err != nil {
		_ = render.Render(w, r, common_http.ErrorInternal(err))
		return;
	}

	render.Respond(w, r, ProductView{
		string(product.ID()),
		product.Name(),
		product.Description(),
		PriceViewFromPrice(product.Price()),
	})
}