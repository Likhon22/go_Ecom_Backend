package user

import (
	"net/http"

	"github.com/Likhon22/ecom/rest/middleware"
)

func (h *Handler) UserRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("POST /users", mngr.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /login", mngr.With(http.HandlerFunc(h.Login)))
}
