package main

import (
	"testing"
)

func TestCalcularImposto(t *testing.T) {
	t.Run("Salário menor que 50.000", func(t *testing.T) {
		salario := 40000.0
		resultadoEsperado := 0.0
		resultado := calcularImposto(salario)
		if resultado != resultadoEsperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", resultadoEsperado, resultado)
		}
	})

	t.Run("Salário maior que 50.000", func(t *testing.T) {
		salario := 100000.0
		resultadoEsperado := 100000.0 * 0.17
		resultado := calcularImposto(salario)
		if resultado != resultadoEsperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", resultadoEsperado, resultado)
		}
	})

	t.Run("Salário maior que 150.000", func(t *testing.T) {
		salario := 200000.0
		resultadoEsperado := 200000.0 * 0.27
		resultado := calcularImposto(salario)
		if resultado != resultadoEsperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", resultadoEsperado, resultado)
		}
	})
}
