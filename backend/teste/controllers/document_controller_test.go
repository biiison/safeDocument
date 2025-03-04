package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"safeCpf/controllers"
	"safeCpf/dto"
	"safeCpf/models"
	"testing"
)

// Mock do DocumentService
type MockDocumentService struct {
	mock.Mock
}

func (m *MockDocumentService) CreateDocument(document dto.CreateDocumentDto) error {
	args := m.Called(document)
	return args.Error(0)
}

func (m *MockDocumentService) GetDocuments() ([]models.DocumentModel, error) {
	args := m.Called()
	return args.Get(0).([]models.DocumentModel), args.Error(1)
}

func (m *MockDocumentService) UpdateBlockDocument(numberDocument string) error {
	args := m.Called(numberDocument)
	return args.Error(0)
}

// Teste para CreateDocument
func TestCreateDocument(t *testing.T) {
	mockService := new(MockDocumentService)
	controller := controllers.NewDocumentController(mockService)

	mockService.On("CreateDocument", mock.Anything).Return(nil)

	document := models.DocumentModel{Numero: "12345678901"}

	body, _ := json.Marshal(document)
	req, err := http.NewRequest("POST", "/create/document", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/create/document", controller.CreateDocument).Methods("POST")

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

// Teste para GetCPFs
func TestGetCPFs(t *testing.T) {
	mockService := new(MockDocumentService)
	controller := controllers.NewDocumentController(mockService)

	mockService.On("GetDocuments").Return([]models.DocumentModel{{Numero: "12345678901"}}, nil)

	req, err := http.NewRequest("GET", "/list-documents", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/list-documents", controller.GetDocuments).Methods("GET")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []models.DocumentModel
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "12345678901", response[0].Numero)

	mockService.AssertExpectations(t)
}

// Teste para BlockCPF
func TestBlockCPF(t *testing.T) {
	mockService := new(MockDocumentService)
	controller := controllers.NewDocumentController(mockService)

	mockService.On("UpdateBlockDocument", "12345678901").Return(nil)

	req, err := http.NewRequest("POST", "/document/block/12345678901", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/document/block/{id}", controller.UpdateBlockDocument).Methods("POST")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "CPF 12345678901 foi bloqueado")
	mockService.AssertExpectations(t)
}

// Teste para Status
func TestStatus(t *testing.T) {
	mockService := new(MockDocumentService)
	controller := controllers.NewDocumentController(mockService)

	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/status", controller.Status).Methods("GET")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Servidor OK")

}
