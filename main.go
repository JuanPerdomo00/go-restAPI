package main

import (
	"fmt"
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

	r.HandleFunc("/", routes.HomeHandler)

	fmt.Printf("Server on http://127.0.0.1:%s\n", PORT)
	http.ListenAndServe(":8080", r)
}
