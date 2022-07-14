package data

import "time"

type Product struct {
	ID			int
	Name		string
	Description string
	Price		float32
	SKU			string
	CreatedOn	string
	UpdatedOn	string
	DeletedOn	string
}

var productList = []*Product{
	&Product {
		ID:			1,
		Name:		"Latte",
		Description: "Frothy Milky Coffee",
		Price:		2.45,
		SKU:		"ahr472",
		CreatedOn:	time.Now().UTC().String(),
		UpdatedOn:	time.Now().UTC().String(),
	},
	&Product {
		ID:			2,
		Name:		"Esspresso",
		Description: "Short and Strong Coffee without milk",
		Price:		1.99,
		SKU:		"kju938",
		CreatedOn:	time.Now().UTC().String(),
		UpdatedOn:	time.Now().UTC().String(),
	},
}