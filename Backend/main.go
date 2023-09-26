package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/angeldev96/mail-indexer/backend/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	goPort := os.Getenv("GO_SERVER_PORT")
	frontendPort := os.Getenv("FRONTEND_SERVER_PORT")
	frontendDir := os.Getenv("FRONTEND_DIR")
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	// Define the flag for the server port
	defaultFrontendPort, _ := strconv.Atoi(frontendPort)
	portPtr := flag.Int("port", defaultFrontendPort, "Port on which the server will run")

	flag.Parse()

	// Start the frontend server on the specified port
	go startFrontendServer(*portPtr, frontendDir)

	r := chi.NewRouter()

	// Configure CORS
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
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

	log.Printf("Starting the Go server on port %s...", goPort)
	listenErr := http.ListenAndServe(fmt.Sprintf(":%s", goPort), r)
	if listenErr != nil {
		log.Fatalf("Could not start the HTTP server: %v", listenErr)
	}

}

func startFrontendServer(port int, frontendDir string) {
	cmd := exec.Command("npm", "run", "dev", "--", fmt.Sprintf("--port=%d", port))
	cmd.Dir = frontendDir
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error starting the frontend server: %v", err)
	}
	log.Printf("Frontend server started on port %d...", port)
}
