package main

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d , Nome: %s , Preço: R$%.2f , Categoria: %s\n",
			product.ID, product.Name, product.Price, product.Category)
	}
}

func getById(id int) *Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func main() {
	p1 := Product{ID: 1, Name: "Notebook", Price: 4500.00, Description: "Notebook gamer", Category: "Eletrônicos"}
	p2 := Product{ID: 2, Name: "Cafeteira", Price: 250.00, Description: "Cafeteira elétrica", Category: "Eletrodomésticos"}

	p1.Save()
	p2.Save()

	p1.GetAll()

	found := getById(3)
	if found != nil {
		fmt.Printf("Produto encontrado: %s - R$%.2f\n", found.Name, found.Price)
	} else {
		fmt.Println("Produto não encontrado.")
	}
}
