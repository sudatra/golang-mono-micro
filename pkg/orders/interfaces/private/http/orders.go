package http

import (
	common_http "golang-mono-micro/pkg/common/http"
	"golang-mono-micro/pkg/orders/application"
	"golang-mono-micro/pkg/orders/domain/orders"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ordersResource struct {
	service application.OrdersService
	repository orders.Repository
}

func AddRoutes(router *chi.Mux, service application.OrdersService, repository orders.Repository) {
	resource := ordersResource{
		service,
		repository,
	};
	
	router.Post("/orders/{id}/paid", resource.PostPaid);
}

func (o ordersResource) PostPaid(w http.ResponseWriter, r *http.Request) {
	cmd := application.MarkOrderAsPaidCommand{
		OrderID: orders.ID(chi.URLParam(r, "id")),
	}

	if err := o.service.MarkOrderAsPaid(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrorInternal(err))
	}

	w.WriteHeader(http.StatusNoContent);
}