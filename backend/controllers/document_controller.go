package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"safeCpf/dto"
	"safeCpf/services"
)

// DocumentController é o controlador que lida com as requisições HTTP para Documents
type DocumentController struct {
	CpfService services.DocumentService
}

// NewDocumentController cria um novo DocumentController com injeção de dependência do serviço
func NewDocumentController(cpfService services.DocumentService) *DocumentController {
	return &DocumentController{
		CpfService: cpfService,
	}
}

// CreateDocument cria um novo Document
func (c *DocumentController) CreateDocument(w http.ResponseWriter, r *http.Request) {
	var cpf dto.CreateDocumentDto
	err := json.NewDecoder(r.Body).Decode(&cpf)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.CpfService.CreateDocument(cpf)
	if err != nil {
		if customErr, ok := err.(*services.CustomError); ok {
			http.Error(w, customErr.Message, customErr.StatusCode)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cpf)
}

func (c *DocumentController) GetDocuments(w http.ResponseWriter, r *http.Request) {
	resp, err := c.CpfService.GetDocuments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	RegistrarConsulta()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *DocumentController) UpdateBlockDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberDocument := vars["id"]
	err := c.CpfService.UpdateBlockDocument(numberDocument)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	RegistrarConsulta()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "CPF %s foi bloqueado", numberDocument)
}

func (c *DocumentController) Status(w http.ResponseWriter, r *http.Request) {
	RegistrarConsulta()
	// Retorna as informações do servidor e a quantidade de consultas realizadas
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Servidor OK. Consultas realizadas: %d", consultas)
}

func RegistrarConsulta() {
	consultas++
}

var (
	consultas int
)
