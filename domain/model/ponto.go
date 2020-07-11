package model

import "time"

type Ponto struct {
	Id_ponto   int32
	DataHora   time.Time
	Id_usuario int32
}
