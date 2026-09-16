package main

import (
	f "fmt"
	// t "time"
)

func main() {

	numero := 5

	// Tabuada do numero
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		f.Printf("%d x %d = %d \n", numero, i, resultado)
	}

	// Loop infinito
	// for {
	// 	f.Println("Golang")
	// 	t.Sleep(2 * time.Second)
	// }

	// For range

	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}

	for _, fruta := range frutas {
		f.Println(fruta)
	}

	for i, fruta := range frutas {
		f.Println("Fruta", fruta, "Indice", i)
	}
}
