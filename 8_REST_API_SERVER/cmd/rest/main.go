package main

import (
	db "Rest_API_Server/internal/database"
	handlers "Rest_API_Server/internal/rest/service"
	"log"
	"net/http"
)

func main() {
	db.Connect()

	http.HandleFunc("GET /", handlers.IndexHandler)
	http.HandleFunc("POST /employees", handlers.AddEmployeeHandler)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}
