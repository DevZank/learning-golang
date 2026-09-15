package main

import "fmt"

func main() {
	fmt.Println(soma(10, 30))

	sum := soma(20, 50)
	fmt.Println(sum)

	sub := soma(80, 20)
	fmt.Println(sub)

	fmt.Println(printName("Oscar", "Mauricio"))

	print1, _ := printText("Hello") // Underscore (_) "burla" o pedido do go da variavel ser usada
	fmt.Println(print1)
}

// Funções que começam com letra Maiúscula são publicas! Podem ser usadas em outros pacotes! main.Soma(1, 5)

func Soma(x int, y int) int {
	return x + y
}

// Funções que começam com letra minúsculas são privadas! Privadas só podem ser usadas no proprio pacote!

func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func printName(name, lastName string) string {
	return ("Olá, " + name + " " + lastName + "!")
}

func printText(text string) (string, string) {
	return text, text
}
