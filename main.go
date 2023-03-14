package main

import (
	"fmt"

	"github.com/Natannms/go_api/models"
	"github.com/Natannms/go_api/routes"
)

func main() {
	ListaDeUsuarios := []models.User{
		{Id: 1, Name: "Natan", Email: "lara@gmail.com", Password: "123"},
		{Id: 2, Name: "Lara", Email: "natan@gmail.com", Password: "123"},
	}

	models.Users = ListaDeUsuarios

	fmt.Println("Running in port http://localhost:8000") // Mensagem com link para executar no nvegador
	routes.HandleRequest()
}
