package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID			int			`json:"id"`
	Name		string		`json:"name"`
	Description string		`json:"description"`
	Price		float32		`json:"price"`
	SKU			string		`json:"sku"`
	CreatedOn	string		`json:"-"`	 // The "-" tag removes the struct property from json output
	UpdatedOn	string		`json:"-"`
	DeletedOn	string		`json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// Calling a slice of products to a http request
func GetProducts() Products {
	return productList
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
		Name:		"Espresso",
		Description: "Short and Strong Coffee without milk",
		Price:		1.99,
		SKU:		"kju938",
		CreatedOn:	time.Now().UTC().String(),
		UpdatedOn:	time.Now().UTC().String(),
	},
}