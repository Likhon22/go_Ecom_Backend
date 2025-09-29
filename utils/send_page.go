package utils

import (
	"net/http"
)

type PaginatedData[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

func SendPage[T any](w http.ResponseWriter, page int64, data []T, count int64, limit int64) {
	paginationData := PaginatedData[T]{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: count,
			TotalPages: (count + limit - 1) / limit,
		},
	}
	SendData(w, paginationData, http.StatusOK)
}
