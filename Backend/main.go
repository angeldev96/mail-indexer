package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/angeldev96/mail-indexer/backend/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	// Define the flag for the server port
	portPtr := flag.Int("port", 8000, "Port on which the server will run")
	flag.Parse()

	// Start the frontend server on the specified port
	go startFrontendServer(*portPtr)

	r := chi.NewRouter()

	// Configure CORS
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:*", "http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsConfig.Handler)

	// Define routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/search", handlers.SearchEmails)
	})

	// Default port for the Go server
	goPort := 8080
	log.Printf("Starting the Go server on port %d...", goPort)
	http.ListenAndServe(fmt.Sprintf(":%d", goPort), r)
}

func startFrontendServer(port int) {
	cmd := exec.Command("npm", "run", "dev", "--", fmt.Sprintf("--port=%d", port))
	cmd.Dir = "/home/luzbel/mail-indexer/Frontend" // Replace with the path to your Vue project
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error starting the frontend server: %v", err)
	}
	log.Printf("Frontend server started on port %d...", port)
}
