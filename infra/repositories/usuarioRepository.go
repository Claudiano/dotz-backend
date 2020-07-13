package repositories

import (
	"log"

	"github.com/claudiano/dotz-backend/domain/model"
	"github.com/claudiano/dotz-backend/infra/database"
)

type UsuarioRepository struct {
}

func (r *UsuarioRepository) FindOne(idUser int32) (model.Usuario, error) {
	db := database.ConnectDB()
	var user model.Usuario

	err := db.Where("idUsuario = ?", idUser).First(&user).Error
	if err != nil {
		log.Printf("Error ao buscar usuario: %v\n", err)
	}

	return user, err
}

func (r *UsuarioRepository) Save(user model.Usuario) (model.Usuario, error) {
	db := database.ConnectDB()
	err := db.Create(&user).Error
	return user, err
}

func (r *UsuarioRepository) FindByEmailAndPassword(email string, password string) (model.Usuario, error) {
	db := database.ConnectDB()
	var user model.Usuario

	err := db.Where("email = ? and password = ?", email, password).First(&user).Error
	if err != nil {
		log.Printf("Error ao buscar usuario: %v\n", err)
	}

	return user, err
}
