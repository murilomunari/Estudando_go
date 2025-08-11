package main

import (
	"fmt"
	"curso-go/model"
	"time"
)


// estudando structs
func main() {
	//criando uma "instancia" da struct Pessoa
	p := model.Pessoa {
		Nome: "Murilo",
		Sobrenome: "Munari",
		Idade: 24,
		Altura: 1.75,
		Endereco: model.Endereco {
			Rua: "Rua dos Bobos",
			Numero: 0,
			Cidade: "São Paulo",
		},
		Emprego: model.Emprego {
			Empresa: "Empresa X",
			Cargo: "Desenvolvedor golang",
			Salario: 10000,
			DataInicio: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	fmt.Println(p)
}

// vendo encapsulamento
