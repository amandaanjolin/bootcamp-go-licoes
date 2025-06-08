package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("---- Dados do Funcionário")
	fmt.Printf("ID (funcionário): %d\n", e.ID)
	fmt.Printf("Nome: %s\n", e.Person.Name)
	fmt.Printf("ID (pessoal): %d\n", e.Person.ID)
	fmt.Printf("Data de nascimento: %s\n", e.Person.DateOfBirth.Format("02/01/2006"))
	fmt.Printf("Cargo: %s\n", e.Position)
}

func main() {
	person := Person{
		ID:          1,
		Name:        "Ana Oliveira",
		DateOfBirth: time.Date(1990, time.May, 12, 0, 0, 0, 0, time.UTC),
	}

	employee := Employee{
		ID:       23,
		Position: "Analista de Sistemas",
		Person:   person,
	}

	employee.PrintEmployee()
}
