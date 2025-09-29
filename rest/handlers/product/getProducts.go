package product

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var count int64
	var wg sync.WaitGroup
	reqQuery := r.URL.Query()
	page, err := strconv.ParseInt(reqQuery.Get("page"), 10, 32)
	if err != nil {
		page = 1
	}
	limit, err := strconv.ParseInt(reqQuery.Get("limit"), 10, 32)
	if err != nil {
		limit = 10
	}

	products, err := h.service.GetAll(page, limit)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		count, err = h.service.Count()
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}
	}()
	wg.Wait()

	utils.SendPage(w, page, products, count, limit)

}
