package main

import (
	"fmt"

	"github.com/claudiano/dotz-backend/domain/model"
	"github.com/claudiano/dotz-backend/infra/repositories"
)

func main() {
	fmt.Println("hello world")

	u := model.Usuario{Nome: "claudiano", Email: "claudiano2@gmail.com", Password: "senha"}

	r := repositories.UsuarioRepository{}
	user, err := r.Save(u)
	if err != nil {
		fmt.Printf("erro encotrado: %v\n", err)
	}

	fmt.Println("Usuario salvo!")
	fmt.Println("usuario: ", user)

}
