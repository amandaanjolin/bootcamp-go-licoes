package main

import (
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func comidaDog(qtd int) float64 {
	return float64(qtd) * 10.0
}

func comidaCat(qtd int) float64 {
	return float64(qtd) * 5.0
}

func comidaHamster(qtd int) float64 {
	return float64(qtd*25) / 100
}

func comidaTarantula(qtd int) float64 {
	return float64(qtd*15) / 100
}

func animal(tipo string) (func(int) float64, string) {
	switch tipo {
	case dog:
		return comidaDog, ""
	case cat:
		return comidaCat, ""
	case hamster:
		return comidaHamster, ""
	case tarantula:
		return comidaTarantula, ""
	default:
		return nil, "animal não reconhecido"
	}
}

func main() {
	animalDog, _ := animal(dog)
	animalCat, _ := animal(cat)
	animalTarantula, _ := animal(tarantula)

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalTarantula(5)

	fmt.Printf("Total de comida necessária: %.2f kg\n", amount)
}
