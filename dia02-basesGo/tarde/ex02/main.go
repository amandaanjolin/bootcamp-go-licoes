package main

import "fmt"

func main() {
	idade := 20
	empregado := true
	anosEmprego := 2
	salario := 120000.0

	if idade <= 22 || !empregado || anosEmprego <= 1 {
		fmt.Println("Empréstimo negado: critérios mínimos não atendidos.")
	} else if salario > 100000 {
		fmt.Println("Empréstimo aprovado: sem juros.")
	} else {
		fmt.Println("Empréstimo aprovado: com juros.")
	}
}
