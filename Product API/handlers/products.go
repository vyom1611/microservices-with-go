package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		p.l.Println("PUT", r.URL.Path)

		//Expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)

		//Finding bad URIs (which are almost same) and doing some validation
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid: More than one ID")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid: More than one capture group")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid: Unable to convert to number")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, w http.ResponseWriter, r *http.Request)
		return
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

func (p *Products) updateProducts(id int,w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUTs")

	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w,"Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
