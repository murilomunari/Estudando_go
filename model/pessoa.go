package model

import "time"

type Pessoa struct {
	Nome string
	Sobrenome string
	Idade int
	Altura float64
	DataDeNascimento time.Time
	Endereco Endereco
	Emprego Emprego
}

func (p Pessoa) IdadeAtual() int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}

// diferemças entre metodo e função

/*
func CalcularIdade(p Pessoa) int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}
*/
