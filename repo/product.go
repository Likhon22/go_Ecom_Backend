package repo

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	Image       string    `json:"image" db:"image"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ProductRepo interface {
	GetAll() ([]*Product, error)
	GetByID(id int) (*Product, error)
	Create(p Product) (*Product, error)
	Update(p Product) (*Product, error)
	Delete(id int) (bool, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}

	return repo
}

func (pr *productRepo) GetAll() ([]*Product, error) {

	var products []*Product
	query := `SELECT id, title, description, price, image FROM products`
	err := pr.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil

}
func (pr *productRepo) GetByID(id int) (*Product, error) {
	var searchedProduct *Product
	query := `SELECT id, title, description, price, image FROM products WHERE id=$1`
	err := pr.db.Get(&searchedProduct, query, id)
	if err != nil {
		fmt.Println("Error fetching product by ID:", err)
		return nil, err
	}

	return searchedProduct, nil
}
func (pr *productRepo) Create(p Product) (*Product, error) {
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
func (pr *productRepo) Update(p Product) (*Product, error) {
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
