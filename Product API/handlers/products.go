package handlers

import (
	"log"
	"net/http"

	"main.go/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//Defining CRUD methods for API
//ServeHTTP is the main entry point for handler
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// GET
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	//POST
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	//PUT
	if r.Method == http.MethodPut {
		//Expect the id in the URI
		p := r.URL.Path

	}

	//Catch all - if not methods is defined
	w.WriteHeader(http.StatusMethodNotAllowed)

}

//getProducts: returns the products from the data store
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {

	//Fetching the data
	lp := data.GetProducts()

	//Serializing data to json
	err := lp.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to convert to json", http.StatusInternalServerError)
	}
}

//addProducts:
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POSTs")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable to handle json", http.StatusBadRequest)
	}

	p.l.Println("Product %#v", prod)
	data.AddProduct(prod)
}
