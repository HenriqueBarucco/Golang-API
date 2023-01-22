package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henriquebarucco/Golang-API/database"
	"github.com/henriquebarucco/Golang-API/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}
