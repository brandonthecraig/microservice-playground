package handlers

import (
	"log"
	"microservice-playground/pkg/data"
	"net/http"
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
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall JSON", http.StatusInternalServerError)
		return
	}
}
