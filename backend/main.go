package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"notas/database"
	"notas/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.InitDb() 
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	router := mux.NewRouter()
	// Public
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/register", routes.RegisterHandler)
	router.HandleFunc("/login", routes.LoginHandler)

	// Private
	router.HandleFunc("/new-note", routes.NewNote)


	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      router,
	}

	log.Printf("Server running on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}