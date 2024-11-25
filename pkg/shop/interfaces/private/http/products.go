package http

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	common_http "golang-mono-micro/pkg/common/http"
	products_domain "golang-mono-micro/pkg/shop/domain"
)

type productsResource struct {
	repo products_domain.Repository
}

type PriceView struct {
	Cents uint `json:"cents"`
	Currency string `json:"currency"`
}

func AddRoutes(router *chi.Mux, repo products_domain.Repository) {
	resource := productsResource{repo};
	router.Get("/products/{id}", resource.Get)
}

func PriceViewFromPrice(p PriceView) PriceView {
	return PriceView{
		p.Cents,
		p.Currency,
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