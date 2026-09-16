package main

import "fmt"

// Structs

/*
	Forma de criar sua própria estrutura de dados
	Personalizar de acordo com a sua necessidade
	Podemos usar vários tipos diferentes
*/

// type <NOME DA ESTRUTURA> struct { <campos> }
type Pessoa struct {
	Nome          string
	Idade         int
	Nacionalidade string
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Isaac", 20, "Brasileiro"})
	fmt.Println(Pessoa{Nome: "Ayrton", Idade: 66, Nacionalidade: "Brasileiro"})
	fmt.Println(Pessoa{Nome: "Rodrigo"})

	p1 := Pessoa{}
	p1.Idade = 30
	p1.Nome = "Samuel"
	p1.Nacionalidade = "Brasileiro"

	fmt.Println(p1)

	p2 := Pessoa{Nome: "Patrick", Idade: 43, Nacionalidade: "Americano"}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)

	fmt.Println(pessoas)

	// Structs + Map
	alunos := map[string][]Pessoa{}
	alunos["Programação"] = pessoas
	fmt.Println(alunos)

	// Outro exemplo
	var alunos2 = map[string][]Pessoa{
		"Programação": {{Nome: "Isaac", Idade: 20}, {Nome: "Bob", Idade: 32}},
		"Engenharia":  {{Nome: "Patrick", Idade: 43}, {Nome: "Teresa", Idade: 40}},
	}
	fmt.Println(alunos2)

	// Struct herdando campos de outra struct 
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}
