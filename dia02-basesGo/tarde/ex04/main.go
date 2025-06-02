package main

import (
	"fmt"
)

func main() {
	employees := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}
	fmt.Println("Idade de Benjamin:", employees["Benjamin"])

	idadeMaior21 := 0
	for _, idade := range employees {
		if idade > 21 {
			idadeMaior21++
		}
	}
	fmt.Println("Funcionários com mais de 21 anos:", idadeMaior21)

	employees["Federico"] = 25
	fmt.Println("Adicionado Federico com 25 anos.")

	delete(employees, "Pedro")
	fmt.Println("Removido Pedro do mapa.")

	fmt.Println("Lista final de funcionários:")
	for nome, idade := range employees {
		fmt.Printf("- %s: %d anos\n", nome, idade)
	}
}
