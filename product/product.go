package product

type Product struct {
	ID          int `json:"id"`
	Title       string
	Description string
	Price       float64
	Image       string
}

type DeleteRequest struct {
	ID int `json:"id"`
}


