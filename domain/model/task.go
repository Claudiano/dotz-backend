package model

import "time"

type Task struct {
	IDTask                int32  `gorm:"AUTO_INCREMENT"`
	Titulo                string `gorm:"type:varchar(200); not null"`
	Status                string `gorm:"type:varchar(50)"`
	DataCriacao           time.Time
	Descricao             string  `gorm:"type:varchar(500)"`
	Estimativa            float64 `gorm:"type:numeric(6,2)"`
	DataUltimaAtualizacao time.Time
	QuantidadeHoras       time.Time

	IDUsuario uint32
}
