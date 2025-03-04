package routes

import (
	"github.com/gorilla/mux"
	"safeCpf/config"
	"safeCpf/controllers"
	"safeCpf/repositories"
	"safeCpf/services"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(r *mux.Router) {
	collection := config.GetCollection("documents")
	validatorRepository := repositories.NewDocumentRepository(collection)
	validatorService := services.NewDocumentService(validatorRepository)
	validatorController := controllers.NewDocumentController(validatorService)

	r.HandleFunc("/create/document", validatorController.CreateDocument).Methods("POST")
	r.HandleFunc("/list-documents", validatorController.GetDocuments).Methods("GET")
	r.HandleFunc("/document/{id}/block", validatorController.UpdateBlockDocument).Methods("PUT")
	r.HandleFunc("/status", validatorController.Status).Methods("GET")
}
