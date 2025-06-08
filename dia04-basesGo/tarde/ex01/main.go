package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name    string
	Surname string
	DNI     int
	Date    time.Time
}

func (s Student) Detail() {
	fmt.Println("----- Detalhes do Aluno -----")
	fmt.Printf("Nome: %s\n", s.Name)
	fmt.Printf("Sobrenome: %s\n", s.Surname)
	fmt.Printf("ID: %d\n", s.DNI)
	fmt.Printf("Data: %s\n", s.Date.Format("02/01/2006"))
}

func main() {
	student := Student{
		Name:    "Carlos",
		Surname: "Silva",
		DNI:     20231234,
		Date:    time.Date(2023, time.February, 14, 0, 0, 0, 0, time.UTC),
	}

	student.Detail()
}
