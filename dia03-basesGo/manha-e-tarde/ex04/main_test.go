package main

import (
	"testing"
)

func TestOperation(t *testing.T) {
	valores := []float64{2, 3, 3, 4, 10, 2, 4, 5}

	t.Run("Teste mínimo", func(t *testing.T) {
		minFunc, err := operation(minimum)
		if err != nil {
			t.Fatalf("Erro inesperado: %v", err)
		}
		resultado := minFunc(valores...)
		esperado := 2.0
		if resultado != esperado {
			t.Errorf("Esperado mínimo %.2f, mas obteve %.2f", esperado, resultado)
		}
	})

	t.Run("Teste máximo", func(t *testing.T) {
		maxFunc, err := operation(maximum)
		if err != nil {
			t.Fatalf("Erro inesperado: %v", err)
		}
		resultado := maxFunc(valores...)
		esperado := 10.0
		if resultado != esperado {
			t.Errorf("Esperado máximo %.2f, mas obteve %.2f", esperado, resultado)
		}
	})

	t.Run("Teste média", func(t *testing.T) {
		avgFunc, err := operation(average)
		if err != nil {
			t.Fatalf("Erro inesperado: %v", err)
		}
		resultado := avgFunc(valores...)
		esperado := 4.125
		if resultado != esperado {
			t.Errorf("Esperado média %.3f, mas obteve %.3f", esperado, resultado)
		}
	})
}
