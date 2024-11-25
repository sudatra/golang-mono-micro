package http

import (
	"net/http"
	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error error `json:"-"`
	HttpStatusCode int `json:"-"`
	AppCode int64 `json:"code,omitempty"`
	ErrorText string `json:"error,omitempty"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode);
	return nil
}

func ErrorInternal(err error) render.Renderer {
	return &ErrorResponse{
		Error: err,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorText: err.Error(),
	};
}

func ErrorBadRequest(err error) render.Renderer {
	return &ErrorResponse{
		Error: err,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorText: err.Error(),
	};
}