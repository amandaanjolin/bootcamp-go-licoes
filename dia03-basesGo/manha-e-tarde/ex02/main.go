package main

import (
	"errors"
	"fmt"
)

func calcularMedia(notas ...int) (float64, error) {
	if len(notas) == 0 {
		return 0, errors.New("nenhuma nota foi fornecida")
	}

	soma := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("não é permitido inserir notas negativas")
		}
		soma += nota
	}

	media := float64(soma) / float64(len(notas))
	return media, nil
}

func main() {
	media, err := calcularMedia(8, 7, 9, 10, 6)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Média do aluno: %.2f\n", media)
	}
}
