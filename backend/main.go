package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"safeCpf/config"
	"safeCpf/routes"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Permite o frontend acessar o backend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Envolve o roteador com o middleware CORS
	handler := corsHandler.Handler(r)

	log.Println("Servidor iniciado na porta 8080...")
	http.ListenAndServe(":8080", handler) // Aqui usamos o handler com CORS
}
