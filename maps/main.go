package main

import "fmt"

func main(){
	idade := map[string]int{}
	idade["isaac"] = 20
	idade["josé"] = 23
	fmt.Println(idade)
	fmt.Println(idade["isaac"])
	fmt.Println(idade["josé"])

	anoNasc := map[string]int{
		"isaac": 2006,
		"josé": 2003,
		"zariel": 2008,
		"arlindo": 2001,
	}
	fmt.Println(anoNasc) // Mostra em ordem alfabética independente da ordem da criação
	fmt.Println(anoNasc["isaac"])
	fmt.Println(anoNasc["josé"])
}
