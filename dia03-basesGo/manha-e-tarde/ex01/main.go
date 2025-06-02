package main

import (
	"fmt"
)

func calcularImposto(salario float64) float64 {
	if salario > 150000 {
		return salario * 0.27
	} else if salario > 50000 {
		return salario * 0.17
	}
	return 0
}

func main() {
	salarios := []float64{40000, 80000, 200000}

	for _, s := range salarios {
		imposto := calcularImposto(s)
		fmt.Printf("Salário: $%.2f | Imposto: $%.2f\n", s, imposto)
	}
}
