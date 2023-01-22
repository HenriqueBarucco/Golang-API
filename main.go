package main

import (
	"fmt"

	"github.com/henriquebarucco/Golang-API/database"
	"github.com/henriquebarucco/Golang-API/models"
	"github.com/henriquebarucco/Golang-API/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
