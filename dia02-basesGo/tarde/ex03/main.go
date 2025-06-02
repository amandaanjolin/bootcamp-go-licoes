package main

import "fmt"

func main() {
	//solucao 1
	// mes := 7

	// meses := []string{
	// 	"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
	// 	"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	// }

	// if mes >= 1 && mes <= 12 {
	// 	fmt.Println(meses[mes-1])
	// } else {
	// 	fmt.Println("Número de mês inválido.")
	// }

	//solucao 2
	mes := 7 // julho

	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Número de mês inválido.")
	}
}
