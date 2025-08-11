package main

import "fmt"

// estudando structs

type Pessoa struct {
	nome string
	sobrenome string
	idade int
	altura float64


}

func main() {
	//criando uma "instancia" da struct Pessoa

	p := Pessoa { // p é uma variavel do tipo Pessoa
		nome: "Murilo",
		sobrenome: "Silva",
		idade: 20,
		altura: 1.75,
	}

	fmt.Println(p)
}
