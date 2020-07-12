package model

type Usuario struct {
	IDUsuario int32  `gorm:"AUTO_INCREMENT"`
	Email     string `gorm:"type:varchar(200);unique"`
	Password  string `gorm:"type:varchar(300);not null"`
	Nome      string `gorm:"type:varchar(150);not null"`

	Pontos []Ponto `gorm:"foreignkey:IDUsuario"`
	Tasks  []Task  `gorm:"foreignkey:IDUsuario"`
}
