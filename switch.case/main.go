package main

import (
	f "fmt"
)

func main() {
	posicao := 2

	switch posicao {
	case 1:
		f.Println("Primeiro Lugar!")
	case 2:
		f.Println("Segundo Lugar!")
	case 3:
		f.Println("Terceiro Lugar!")
	}
}
