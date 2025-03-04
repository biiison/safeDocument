package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"safeCpf/models"
	"safeCpf/services"
	"testing"
)

type MockDocumentRepository struct {
	mock.Mock
}

func (m *MockDocumentRepository) CreateDocument(document models.DocumentModel) error {
	args := m.Called(document)
	return args.Error(0)
}

func (m *MockDocumentRepository) GetDocuments() ([]models.DocumentModel, error) {
	args := m.Called()
	return args.Get(0).([]models.DocumentModel), args.Error(1)
}

func (m *MockDocumentRepository) UpdateBlockDocument(numberDocument string) error {
	args := m.Called(numberDocument)
	return args.Error(0)
}

func TestGetDocuments(t *testing.T) {
	mockRepo := new(MockDocumentRepository)
	service := services.NewDocumentService(mockRepo)

	mockRepo.On("GetDocuments").Return([]models.DocumentModel{
		{Numero: "123.456.789-00", Tipo: "cpf"},
		{Numero: "234.567.890-01", Tipo: "cpf"},
	}, nil)

	docs, err := service.GetDocuments()

	assert.NoError(t, err)
	assert.Len(t, docs, 2)
	assert.Equal(t, "123.456.789-00", docs[0].Numero)
	mockRepo.AssertExpectations(t)
}

func TestUpdateBlockDocument(t *testing.T) {
	mockRepo := new(MockDocumentRepository)
	service := services.NewDocumentService(mockRepo)

	tests := []struct {
		name        string
		documentID  string
		expectedErr error
		mockReturn  error
	}{
		{
			name:        "Successfully block document",
			documentID:  "123.456.789-00",
			expectedErr: nil,
			mockReturn:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("UpdateBlockDocument", tt.documentID).Return(tt.mockReturn)

			err := service.UpdateBlockDocument(tt.documentID)

			if tt.expectedErr != nil {
				// Verifica se o erro contém a mensagem esperada
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error()) // Aqui usamos assert.Contains
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
