package database

import "github.com/Likhon22/ecom/product"

func init() {
	product1 := product.Product{
		ID:          1,
		Title:       "Product 1",
		Description: "Description 1",
		Price:       100,
		Image:       "Image 1",
	}
	product2 := product.Product{
		ID:          2,
		Title:       "Product 2",
		Description: "Description 2",
		Price:       200,
		Image:       "Image 2",
	}
	product3 := product.Product{
		ID:          3,
		Title:       "Product 3",
		Description: "Description 3",
		Price:       300,
		Image:       "Image 3",
	}
	product.ProductList = append(product.ProductList, product1, product2, product3)

}
