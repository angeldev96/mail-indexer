package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/angeldev96/mail-indexer/backend/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func loadEnvVariables() error {
	return godotenv.Load("../.env")
}

func configureCORS(allowedOrigins []string) func(next http.Handler) http.Handler {
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	return corsConfig.Handler
}

func startGoServer(port string) {
	r := chi.NewRouter()
	r.Use(configureCORS(strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")))
	r.Route("/api", func(r chi.Router) {
		r.Get("/search", handlers.SearchRecords)
	})

	log.Printf("Starting the Go server on port %s...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalf("Could not start the HTTP server: %v", err)
	}
}

func main() {
	if err := loadEnvVariables(); err != nil {
		log.Fatal("Error loading .env file")
	}

	goPort := os.Getenv("GO_SERVER_PORT")
	frontendPort := os.Getenv("FRONTEND_SERVER_PORT")
	frontendDir := os.Getenv("FRONTEND_DIR")

	defaultFrontendPort, _ := strconv.Atoi(frontendPort)
	portPtr := flag.Int("port", defaultFrontendPort, "Port on which the server will run")
	flag.Parse()

	go startFrontendServer(*portPtr, frontendDir)
	startGoServer(goPort)
}

func startFrontendServer(port int, frontendDir string) {
	distDir := filepath.Join(frontendDir, "dist")

	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		// If the dist directory does not exist, build the project
		buildCmd := exec.Command("npm", "run", "build")
		buildCmd.Dir = frontendDir
		err := buildCmd.Run()
		if err != nil {
			log.Fatalf("Error building the frontend project: %v", err)
		}
	}

	// Start the preview server
	previewCmd := exec.Command("npm", "run", "preview", "--", "--port", strconv.Itoa(port))
	previewCmd.Dir = frontendDir
	err := previewCmd.Start()
	if err != nil {
		log.Fatalf("Error starting the preview server: %v", err)
	}
	log.Printf("Preview server started on port %d...", port)
}
