package repositories

import (
	"github.com/claudiano/dotz-backend/domain/model"
	"github.com/claudiano/dotz-backend/infra/database"
)

type UsuarioRepository struct {
}

func (r *UsuarioRepository) FindOne(user model.Usuario) (model.Usuario, error) {

	return model.Usuario{}, nil
}

func (r *UsuarioRepository) Save(user model.Usuario) (model.Usuario, error) {
	db := database.ConnectDB()
	err := db.Create(&user).Error
	return user, err
}
