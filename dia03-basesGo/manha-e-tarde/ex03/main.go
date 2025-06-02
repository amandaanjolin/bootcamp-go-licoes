package main

import (
	"errors"
	"fmt"
)

func calcularSalario(minutosTrabalhados int, categoria string) (float64, error) {
	horas := float64(minutosTrabalhados) / 60

	var salarioBase float64
	var bonusPercentual float64

	switch categoria {
	case "C":
		salarioBase = horas * 1000
	case "B":
		salarioBase = horas * 1500
		bonusPercentual = 0.20
	case "A":
		salarioBase = horas * 3000
		bonusPercentual = 0.50
	default:
		return 0, errors.New("categoria inválida")
	}

	salarioFinal := salarioBase + (salarioBase * bonusPercentual)
	return salarioFinal, nil
}

func main() {
	minutos := 9600
	categoria := "A"

	salario, err := calcularSalario(minutos, categoria)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Salário para categoria %s: $%.2f\n", categoria, salario)
	}
}
