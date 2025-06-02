package main

import (
	"testing"
)

func TestCalcularMedia(t *testing.T) {
	t.Run("Média de notas válidas", func(t *testing.T) {
		media, err := calcularMedia(8, 7, 6)
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		esperado := 7.0
		if media != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, media)
		}
	})

	t.Run("Notas negativas não são permitidas", func(t *testing.T) {
		_, err := calcularMedia(10, -2, 8)
		if err == nil {
			t.Errorf("Esperava erro para nota negativa, mas não ocorreu")
		}
	})

	t.Run("Nenhuma nota fornecida", func(t *testing.T) {
		_, err := calcularMedia()
		if err == nil {
			t.Errorf("Esperava erro por ausência de notas, mas não ocorreu")
		}
	})
}
