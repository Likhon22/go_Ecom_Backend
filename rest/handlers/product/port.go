package product

import "github.com/Likhon22/ecom/domain"

type ProductService interface {
	GetAll(page, limit int64) ([]*domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(p domain.Product) (*domain.Product, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(id int) (bool, error)
	Count() (int64, error)
}
