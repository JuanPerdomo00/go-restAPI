package main

import (
	"log"
	"net/http"

	"github.com/JuanPerdomo00/go-restAPI/db"
	"github.com/JuanPerdomo00/go-restAPI/models"
	"github.com/JuanPerdomo00/go-restAPI/routes"
	"github.com/gorilla/mux"
)

func main() {
	const PORT = "8080"
	db.DBconnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler).Methods("GET")

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")

	log.Printf("Server on http://127.0.0.1:%s\n", PORT)
	http.ListenAndServe(":8080", r)
}
