package muxrouters

import (
	"github.com/JuanPerdomo00/go-restAPI/routes"
	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func HandlerUserMain() {
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")
}
