package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

var products []Product

func loadProducts(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &products)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}

func searchProductsByPrice(w http.ResponseWriter, r *http.Request) {
	priceStr := r.URL.Query().Get("priceGt")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "invalid priceGt parameter", http.StatusBadRequest)
		return
	}

	var filtered []Product
	for _, p := range products {
		if p.Price > price {
			filtered = append(filtered, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func main() {
	if err := loadProducts("products.json"); err != nil {
		fmt.Println("Erro ao carregar produtos:", err)
		return
	}

	r := chi.NewRouter()

	r.Get("/ping", pingHandler)
	r.Get("/products", getAllProducts)
	r.Get("/products/{id}", getProductByID)
	r.Get("/products/search", searchProductsByPrice)

	http.ListenAndServe(":8080", r)
}
