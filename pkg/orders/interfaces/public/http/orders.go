package http

import (
	common_http "golang-mono-micro/pkg/common/http"
	"golang-mono-micro/pkg/orders/application"
	"golang-mono-micro/pkg/orders/domain/orders"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ordersResource struct {
	service application.OrdersService
	repository orders.Repository
}

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address PostOrderAddress `json:"address"`
}

type PostOrderAddress struct {
	Name string `json:"name"`
	Street string `json:"street"`
	City string `json:"city"`
	PostCode string `json:"post_code"`
	Country string `json:"country"`
}

type PostOrdersResponse struct {
	OrderID string
}

type OrderPaidView struct {
	ID string `json:"id"`
	IsPaid bool `json:"is_paid"`
}

func AddRoutes(router *chi.Mux, service application.OrdersService, repository orders.Repository) {
	resource := ordersResource{
		service,
		repository,
	}

	router.Post("/orders", resource.Post);
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{};
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, common_http.ErrorBadRequest(err));
		return;
	}

	cmd := application.PlaceOrderCommand{
		OrderID: orders.ID(uuid.NewV1().String()),
		ProductID: req.ProductID,
		Address: application.PlaceOrderCommandAddress(req.Address),
	}

	if err := o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrorInternal(err));
		return;
	}

	w.WriteHeader(http.StatusOK);
	render.JSON(w, r, PostOrdersResponse{
		OrderID: string(cmd.OrderID),
	})
}

func (o ordersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	order, err := o.repository.ByID(orders.ID(chi.URLParam(r, "id")));
	if err != nil {
		_ = render.Render(w, r, common_http.ErrorBadRequest(err));
		return;
	}

	render.Respond(w, r, OrderPaidView{
		string(order.ID()),
		order.Paid(),
	})
}