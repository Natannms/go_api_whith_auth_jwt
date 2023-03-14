package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Natannms/go_api/database"
	"github.com/Natannms/go_api/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func Home(r http.ResponseWriter, rq *http.Request) {
	fmt.Println("Seja bem vindo ao go_api")
}

func CreateUser(r http.ResponseWriter, rq *http.Request) {
	// var user models.User
	// json.NewDecoder(rq.Body).Decode(&user)
	// database.Conn().Create(&user)
	// json.NewEncoder(r).Encode(user)

	var user models.User
	json.NewDecoder(rq.Body).Decode(&user)

	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	database.Conn().Create(&user)
	json.NewEncoder(r).Encode(user)
}

func AllUsers(r http.ResponseWriter, rq *http.Request) {

	var users []models.User
	database.Conn().Find(&users)
	json.NewEncoder(r).Encode(users)
}

func GetUser(r http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	id := vars["id"]
	var user models.User

	result := database.Conn().First(&user, id)
	if result.Error != nil {
		http.Error(r, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(r).Encode(user)
}

func UpdateUser(r http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	id := vars["id"]

	var user models.User

	result := database.Conn().First(&user, id)
	if result.Error != nil {
		http.Error(r, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewDecoder(rq.Body).Decode(&user)
	database.Conn().Save(&user)
	json.NewEncoder(r).Encode(&user)
}

func DeleteUser(r http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	id := vars["id"]

	var user models.User

	result := database.Conn().First(&user, id)
	if result.Error != nil {
		http.Error(r, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	database.Conn().Delete(&user, id)
	json.NewEncoder(r).Encode(&user)
}
