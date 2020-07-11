package main

import (
	"fmt"

	"github.com/claudiano/dotz-backend/domain/model"
)

func main() {
	fmt.Println("hello world")

	u := model.Usuario{Id_usuario: 1, Nome: "claudiano", Email: "claudiano@gmail.com", Password: "senha"}

	fmt.Println(u)
}
