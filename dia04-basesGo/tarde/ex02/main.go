package main

import (
	"fmt"
)

type Product interface {
	Price() float64
}

type SmallProduct struct {
	basePrice float64
}

type MediumProduct struct {
	basePrice float64
}

type LargeProduct struct {
	basePrice float64
}

func (s SmallProduct) Price() float64 {
	return s.basePrice
}

func (m MediumProduct) Price() float64 {
	// preço + 3% + 3%
	return m.basePrice * 1.06
}

func (l LargeProduct) Price() float64 {
	// preço + 6% + frete fixo de 2500
	return l.basePrice*1.06 + 2500
}

func NewProduct(productType string, basePrice float64) Product {
	switch productType {
	case "small":
		return SmallProduct{basePrice}
	case "medium":
		return MediumProduct{basePrice}
	case "large":
		return LargeProduct{basePrice}
	default:
		return nil
	}
}

func main() {
	p1 := NewProduct("small", 1000)
	p2 := NewProduct("medium", 1000)
	p3 := NewProduct("large", 1000)

	fmt.Printf("Preço total (Small): $%.2f\n", p1.Price())
	fmt.Printf("Preço total (Medium): $%.2f\n", p2.Price())
	fmt.Printf("Preço total (Large): $%.2f\n", p3.Price())
}
