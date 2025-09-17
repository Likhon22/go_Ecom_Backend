package repo

type Product struct {
	ID          int `json:"id"`
	Title       string
	Description string
	Price       float64
	Image       string
}

type ProductRepo interface {
	GetAll() ([]*Product, error)
	GetByID(id int) (*Product, error)
	Create(p Product) (*Product, error)
	Update(p Product) (*Product, error)
	Delete(id int) (bool, error)
}

type productRepo struct {
	Products []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateProduct(repo)
	return repo
}

func (pr *productRepo) GetAll() ([]*Product, error) {
	return pr.Products, nil
}
func (pr *productRepo) GetByID(id int) (*Product, error) {
	var searchedProduct *Product
	for _, product := range pr.Products {
		if product.ID == id {
			searchedProduct = product
			break

		}

	}
	return searchedProduct, nil
}
func (pr *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(pr.Products) + 1
	pr.Products = append(pr.Products, &p)
	return &p, nil
}
func (pr *productRepo) Update(p Product) (*Product, error) {
	for i, product := range pr.Products {
		if product.ID == p.ID {
			pr.Products[i] = &p
			return &p, nil
		}
	}
	return nil, nil
}

func (pr *productRepo) Delete(id int) (bool, error) {
	for i, p := range pr.Products {
		if p.ID == id {
			pr.Products = append(pr.Products[:i], pr.Products[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func generateProduct(r *productRepo) {

	product1 := &Product{
		ID:          1,
		Title:       "Product 1",
		Description: "Description 1",
		Price:       100,
		Image:       "Image 1",
	}
	product2 := &Product{
		ID:          2,
		Title:       "Product 2",
		Description: "Description 2",
		Price:       200,
		Image:       "Image 2",
	}
	product3 := &Product{
		ID:          3,
		Title:       "Product 3",
		Description: "Description 3",
		Price:       300,
		Image:       "Image 3",
	}
	r.Products = append(r.Products, product1, product2, product3)

}
