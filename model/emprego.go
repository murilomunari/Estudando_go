package model

import "time"


type Emprego struct {
	Empresa string
	Cargo string
	Salario float64
	DataInicio time.Time
}