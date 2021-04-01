package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID   string
	Name string
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
