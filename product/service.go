package product

import "github.com/Likhon22/ecom/domain"

type productService struct {
	repo ProductRepo
}

func NewProductService(repo ProductRepo) ProductService {
	return &productService{repo: repo}
}

func (ps *productService) GetAll(page, limit int64) ([]*domain.Product, error) {
	proudcts, err := ps.repo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}
	return proudcts, nil
}

func (ps *productService) GetByID(id int) (*domain.Product, error) {
	product, err := ps.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *productService) Create(p domain.Product) (*domain.Product, error) {
	createdProduct, err := ps.repo.Create(p)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}
func (ps *productService) Update(p domain.Product) (*domain.Product, error) {
	updatedProduct, err := ps.repo.Update(p)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}
func (ps *productService) Delete(id int) (bool, error) {
	isDeleted, err := ps.repo.Delete(id)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}

func (ps *productService) Count() (int64, error) {
	count, err := ps.repo.Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
