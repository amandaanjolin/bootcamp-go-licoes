package main

import (
	"fmt"
)

func main() {
	palavra := "teste"

	fmt.Printf("A palavra '%s' tem %d letras\n", palavra, len(palavra))

	for i := 0; i < len(palavra); i++ {
		fmt.Printf("%c\n", palavra[i])
	}
}
