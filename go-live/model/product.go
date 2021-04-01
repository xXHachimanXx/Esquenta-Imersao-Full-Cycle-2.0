package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID   string `json:id`
	Name string `json:name`
}

type Products struct {
	Products []*Product
}

func (products *Products) Add(product *Product) {
	products.Products = append(products.Products, product)
}

func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}
