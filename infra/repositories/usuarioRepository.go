package repositories

import "github.com/claudiano/dotz-backend/domain/model"

type UsuarioRepository struct {
}

func (r *UsuarioRepository) FindOne(user model.Usuario) (model.Usuario, error) {
	return model.Usuario{}, nil
}

func (r *UsuarioRepository) Save(user model.Usuario) (model.Usuario, error) {
	return model.Usuario{}, nil
}
