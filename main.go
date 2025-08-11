package main

import (
	"fmt"
	"curso-go/model"
)


// estudando structs
func main() {
	//criando uma "instancia" da struct Pessoa

	p := model.Pessoa { // p é uma variavel do tipo Pessoa
		Nome: "Murilo",
		Sobrenome: "Silva",
		Idade: 20,
		Altura: 1.75,
	}

	fmt.Println(p)
}

// vendo encapsulamento
