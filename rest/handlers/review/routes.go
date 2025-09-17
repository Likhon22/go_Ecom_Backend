package review

import (
	"net/http"

	"github.com/Likhon22/ecom/rest/middleware"
)

func (h *Handler) ReviewRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /reviews", mngr.With(http.HandlerFunc(h.GetReview)))

}
