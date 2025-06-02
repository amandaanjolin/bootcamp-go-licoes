package main

import "testing"

func TestCalcularSalario(t *testing.T) {
	t.Run("Categoria A", func(t *testing.T) {
		salario, err := calcularSalario(120, "A")
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		esperado := 9000.00
		if salario != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, salario)
		}
	})

	t.Run("Categoria B", func(t *testing.T) {
		salario, err := calcularSalario(180, "B")
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		esperado := 5400.00
		if salario != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, salario)
		}
	})

	t.Run("Categoria C", func(t *testing.T) {
		salario, err := calcularSalario(60, "C")
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		esperado := 1000.00
		if salario != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, salario)
		}
	})
}
