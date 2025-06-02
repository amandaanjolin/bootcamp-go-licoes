package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDog(t *testing.T) {
	esperado := 100.0
	resultado := comidaDog(10)
	if resultado != esperado {
		t.Errorf("Esperado %.2f para cães, mas obteve %.2f", esperado, resultado)
	}
}

func TestCat(t *testing.T) {
	esperado := 50.0
	resultado := comidaCat(10)
	if resultado != esperado {
		t.Errorf("Esperado %.2f para gatos, mas obteve %.2f", esperado, resultado)
	}
}

func TestHamster(t *testing.T) {
	esperado := 0.75
	resultado := comidaHamster(3)
	if resultado != esperado {
		t.Errorf("Esperado %.2f para hamsters, mas obteve %.2f", esperado, resultado)
	}
}

func TestTarantula(t *testing.T) {
	esperado := float64(0.45)
	resultado := comidaTarantula(3)
	require.Equal(t, esperado, resultado)
}
