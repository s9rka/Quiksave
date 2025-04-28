package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"notas/database"
	"notas/middleware"
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

	secretKey := []byte(os.Getenv("SECRET_KEY"))

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/register", routes.RegisterHandler)
	router.HandleFunc("/login", routes.LoginHandler)

	private := router.PathPrefix("/").Subrouter()
	private.HandleFunc("/create-vault", routes.CreateVault)
	private.HandleFunc("/get-vaults", routes.GetVaults)
	private.HandleFunc("/vault/{id:[0-9]+}", routes.GetVaultByID).Methods(http.MethodGet)
	private.HandleFunc("/create-note", routes.CreateNote)
	private.HandleFunc("/get-notes", routes.GetNotes)
	private.HandleFunc("/note/{id:[0-9]+}", routes.GetNoteByID).Methods(http.MethodGet)
	private.HandleFunc("/note/{id:[0-9]+}", routes.DeleteNote).Methods(http.MethodDelete)
	private.HandleFunc("/note/{id:[0-9]+}", routes.EditNote).Methods(http.MethodPut)
	private.HandleFunc("/tags", routes.GetUserTags)
	private.HandleFunc("/logout", routes.Logout)
	private.HandleFunc("/me", routes.GetMe)
	private.HandleFunc("/refresh", routes.RefreshJWT)

	private.Use(middleware.CreateAuthMiddleware(secretKey))

	handler := middleware.LogRequestMiddleware(
		middleware.CORSMiddleware(router),
	)

	// Load your certificate and key files
	certFile := "../localhost+2.pem"
	keyFile := "../localhost+2-key.pem"

	// Configure TLS if you want advanced settings (optional)
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,

		// Customize other TLS settings if needed
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Addr:         addr,
		TLSConfig:    tlsConfig,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      handler,
	}

	log.Printf("Starting HTTPS server on https://localhost:8443")
	// ListenAndServeTLS will load the certificates for HTTPS
	err = srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}
