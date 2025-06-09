package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("error: salary is less than 10000")

func verifySalary(salary int) error {
	if salary <= 10000 {
		return err1
	}
	return nil
}

func main() {
	salary := 9000
	err := verifySalary(salary)
	if errors.Is(err, err1) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Salary is valid")
	}

}
