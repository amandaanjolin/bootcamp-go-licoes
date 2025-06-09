package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("error: the salary entered does not reach the taxable minimum")

func verifySalary(salary int) error {
	if salary < 150000 {
		return err1
	}
	return nil
}

func main() {
	salary := 140000

	err := verifySalary(salary)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Must pay tax")
	}
}
