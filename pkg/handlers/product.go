package handlers

import (
	"errors"
	"fmt"
	"log"
	"microservice-playground/pkg/data"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw, r)
	case http.MethodPost:
		p.addProduct(rw, r)
	case http.MethodPut:
		p.l.Println("handle put request")

		path := r.URL.Path
		fmt.Println(path)
		reg := regexp.MustCompile(`/([0-9]+)`)
		s := reg.FindAllStringSubmatch(path, -1)
		if len(s) != 1 {
			p.l.Println("invalid URL more than 1 id")
			http.Error(rw, "invalid URI", http.StatusBadRequest)
			return
		}

		if len(s[0]) != 2 {
			p.l.Println("invalid URL more than 1 capture group")
			http.Error(rw, "invalid URI", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(s[0][1])
		if err != nil {
			p.l.Println("valid URL unable to convert to number")
			http.Error(rw, "invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle get request")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall JSON", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle post request")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshall JSON", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle put request")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshall JSON", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if errors.Is(err, data.ErrNotFound) {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "error while updating product", http.StatusInternalServerError)
		return
	}

}
