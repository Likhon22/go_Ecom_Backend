package database

import "github.com/Likhon22/ecom/product"

var productList []product.Product

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
	productList = append(productList, product1, product2, product3)
}

func StoreProduct(p product.Product) {
	p.ID = len(productList) + 1
	productList = append(productList, p)
}

func List() []product.Product {
	return productList
}
func GetProductByID(id int) *product.Product {
	var searchedProduct *product.Product
	for _, product := range productList {
		if product.ID == id {
			searchedProduct = &product
			break

		}

	}
	return searchedProduct
}

func UpdateProduct(updatedProduct product.Product) product.Product {
	for i, p := range productList {
		if p.ID == updatedProduct.ID {
			productList[i] = updatedProduct
			return updatedProduct
		}
	}
	return product.Product{}
}

func DeleteProduct(id int) bool {
	for i, p := range productList {
		if p.ID == id {
			productList = append(productList[:i], productList[i+1:]...)
			return true
		}
	}
	return false
}
