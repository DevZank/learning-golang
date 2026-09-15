package main

import "fmt"

func main() {
	// Variavies
	var nome string
	nome = "Isaac"
	fmt.Println(nome)

	nome = "José"
	fmt.Println(nome)

	const passaport = "Aprovado" // Go faz inferência de tipos

	var b, c int = 1, 2
	fmt.Println(b + c)

	var exemplo1 = true
	fmt.Println(exemplo1)

	exemplo2 := true
	fmt.Println(exemplo2)
}
