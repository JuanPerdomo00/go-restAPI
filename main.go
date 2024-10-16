package main

import (
	"log"
	"net/http"

	"github.com/JuanPerdomo00/go-restAPI/db"
	"github.com/JuanPerdomo00/go-restAPI/models"
	muxrouters "github.com/JuanPerdomo00/go-restAPI/mux-routers" // muxrouters as
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	const PORT = "8080"
	db.DBconnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	muxrouters.HandlerUserMain() // Handlers Users

	log.Printf("Server on http://127.0.0.1:%s\n", PORT)
	http.ListenAndServe(":8080", r)
}
