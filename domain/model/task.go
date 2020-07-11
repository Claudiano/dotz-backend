package model

import "time"

type Task struct {
	Id_task               int32
	Titulo                string
	Status                string
	DataCriacao           time.Time
	Descricao             string
	Estimativa            float64
	DataUltimaAtualizacao time.Time
	QuantidadeHoras       time.Time

	Usuario
}
