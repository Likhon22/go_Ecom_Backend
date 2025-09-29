package repo

import (
	"fmt"

	"github.com/Likhon22/ecom/domain"
	"github.com/Likhon22/ecom/product"
	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

type ProductRepo interface {
	product.ProductService
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}

	return repo
}

func (pr *productRepo) GetAll(page, limit int64) ([]*domain.Product, error) {
	offset := (page - 1) * limit
	fmt.Println("Limit:", limit, "Offset:", offset)
	var products []*domain.Product
	query := `SELECT id, title, description, price, image FROM products LIMIT $1 OFFSET $2`
	err := pr.db.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil

}
func (pr *productRepo) GetByID(id int) (*domain.Product, error) {
	var searchedProduct *domain.Product
	query := `SELECT id, title, description, price, image FROM products WHERE id=$1`
	err := pr.db.Get(&searchedProduct, query, id)
	if err != nil {
		fmt.Println("Error fetching product by ID:", err)
		return nil, err
	}

	return searchedProduct, nil
}
func (pr *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	INSERT INTO products (title, description, price, image)
	VALUES (:title, :description, :price, :image)
	RETURNING id
	`
	rows, err := pr.db.NamedQuery(query, p)
	if err != nil {
		fmt.Println("Error inserting product:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&p.ID)
	}
	return &p, nil
}
func (pr *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products 
		SET title=$1, description=$2, price=$3, image=$4, updated_at=NOW() 
		WHERE id=$5
	`
	result, err := pr.db.Exec(query, p.Title, p.Description, p.Price, p.Image, p.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	return &p, nil
}

func (pr *productRepo) Delete(id int) (bool, error) {

	query := `DELETE FROM products WHERE id=$1`
	result, err := pr.db.Exec(query, id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func (pr *productRepo) Count() (int64, error) {
	var count int64
	query := `SELECT COUNT(*) FROM products`
	err := pr.db.Get(&count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
