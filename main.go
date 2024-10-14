package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const PORT = "8080"
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World desde go"))
	})

	fmt.Printf("Server on http://127.0.0.1:%s\n", PORT)
	http.ListenAndServe(":8080", r)
}
