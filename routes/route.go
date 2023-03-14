package routes

import (
	"log"
	"net/http"

	"github.com/Natannms/go_api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/users", controllers.AllUsers)
	r.HandleFunc("/user", controllers.CreateUser)
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
