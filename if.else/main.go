package main

import "fmt"

func main() {
	numero := 1

	if numero == 1 {
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("Valor NÃO é igual a 1!")
	}

	if numero == 1 {
		fmt.Println("Valor é igual a 1")
	} else if numero == 2 {
		fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("O valor é diferente de 1 e 2")
	}

	x := 9

	if x%2 == 0 {
		fmt.Printf("%d é par", x)
	} else {
		fmt.Printf("%d não é par", x)
	}
}
