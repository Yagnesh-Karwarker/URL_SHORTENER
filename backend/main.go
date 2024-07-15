package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshortener/database"
	"urlshortener/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db, err := database.InitDB("./urlshortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	h := handlers.NewHandler(db)

	r.HandleFunc("/api/shorten", h.ShortenURL).Methods("POST")
	r.HandleFunc("/api/{shortURL}", h.RedirectURL).Methods("GET")

	//log.Println("Server started on :8080")
	//log.Fatal(http.ListenAndServe(":8080", r))

	handler := cors.Default().Handler(r)
	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
