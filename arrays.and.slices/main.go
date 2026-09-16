package main

import "fmt"

func main() {

	// Arrays

	// var array [2]string
	// array[0] = "Hello"
	// array[1] = "World"
	// fmt.Println(array[0], array[1])
	// fmt.Println(array)

	// numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(numPrimos)
	// fmt.Println(numPrimos[0:4]) // retorna tudo até a posição 4 mas não retorna a posição 4
	// fmt.Println(numPrimos[5:])  // retorna tudo depois da posição 1 - se inclui (mostra o valor da posição)
	// fmt.Println(numPrimos[:5])  // retorna tudo antes da posição 1 - não se inclui (não mostra o valor da posição)

	// Slices

	// var slice []string
	slice := make([]string, 3)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	slice[2] = "Isaac"
	fmt.Println(slice[2])
	fmt.Println(slice)

	numPares := []int{2, 4, 6, 8}
	fmt.Println(numPares)

	numPares = append(numPares, 10, 12, 14, 16, 18, 20)
	fmt.Println(numPares)

	// LISTAS

	/*
		1 - Arrays e Slices: Homogêneos
			todos os elementos tem o mesmo tipo
		[1, 2, 3, 4, 5, 6] - []int
		["zank", "isaac", "dev"] - []string

		2 - Maps: Heterogêneos
			pode misturar tipos
			estrutura chave - valor
		[key] = value
			chave tem um tipo, e o valor pode ter outra
		map [string]int
			{"isaac": 20, "zank", 20}
		map [string]string
			{"isaac": "zank", "dev": "developer"}

		Array
			Tamanho fixo, de zero ou mais elementos do mesmo tipo
			acessamos os valores com índice: a[0], a[1]...
			função embutida len() retorna o tamanho do array
			por conta do tamanho fixo, não é tão utilizado

		Slice
			Semelhante ao array mas sem tamanho fixo
			acessamos os valores com índice: a[0], a[1]...
			função embutida len() retorna o tamanho do array
			função append() para adicionar valores no slice
	*/
}
