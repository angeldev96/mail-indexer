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
	// Define el flag para el puerto del servidor Vue
	vuePortPtr := flag.Int("port", 8000, "Puerto en el que se ejecutará el servidor Vue")
	flag.Parse()

	// Iniciar el servidor Vue en el puerto especificado
	go startVueServer(*vuePortPtr)

	r := chi.NewRouter()

	// Configura CORS
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:7070", "http://localhost:7070"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsConfig.Handler)

	// Definir rutas
	r.Route("/api", func(r chi.Router) {
		r.Get("/search", handlers.SearchEmails)
	})

	// Puerto predeterminado para el servidor Go
	goPort := 8080
	log.Printf("Iniciando el servidor Go en el puerto %d...", goPort)
	http.ListenAndServe(fmt.Sprintf(":%d", goPort), r)
}

func startVueServer(port int) {
	cmd := exec.Command("npm", "run", "dev", "--", fmt.Sprintf("--port=%d", port))
	cmd.Dir = "/home/luzbel/mail-indexer/frontend" // Reemplaza con la ruta a tu proyecto Vue
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error iniciando el servidor Vue: %v", err)
	}
	log.Printf("Servidor Vue iniciado en el puerto %d...", port)
}
