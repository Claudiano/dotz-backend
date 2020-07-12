package model

import "time"

type Ponto struct {
	IDPonto   int32 `gorm:AUTO_INCREMENT`
	DataHora  time.Time
	IDUsuario int32
}
