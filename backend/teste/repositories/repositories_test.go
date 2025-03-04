package repositories_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"safeCpf/models"
	"testing"
)

// MockRepository simula a interface DocumentRepository
type MockRepository struct {
	mock.Mock
}

// Implementação dos métodos mockados
func (m *MockRepository) CreateDocument(cpf models.DocumentModel) error {
	args := m.Called(cpf)
	return args.Error(0)
}

func (m *MockRepository) GetDocuments() ([]models.DocumentModel, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.DocumentModel), args.Error(1)
}

func (m *MockRepository) UpdateBlockDocument(numberDocument string) error {
	args := m.Called(numberDocument)
	return args.Error(0)
}

// Teste do método GetDocuments
func TestGetDocuments(t *testing.T) {
	// Criando uma instância do mock
	mockRepo := new(MockRepository)

	// Teste de sucesso: configurando o mock para retornar documentos
	mockRepo.On("GetDocuments").Return([]models.DocumentModel{
		{Numero: "123.456.789-00", Tipo: "cpf", Valido: true, Block: false},
	}, nil).Once() // .Once() garante que o mock é chamado apenas uma vez

	// Chamando o método GetDocuments
	documents, err := mockRepo.GetDocuments()

	// Verificando se o erro foi nil e o número de documentos
	assert.NoError(t, err)
	assert.Len(t, documents, 1)
	assert.Equal(t, "123.456.789-00", documents[0].Numero)
	assert.Equal(t, "cpf", documents[0].Tipo)

	// Teste de erro: configurando o mock para retornar um erro
	mockRepo.On("GetDocuments").Return(nil, errors.New("erro ao consultar")).Once()

	// Chamando o método GetDocuments novamente e verificando o erro
	documents, err = mockRepo.GetDocuments()

	// Verificando se o erro foi retornado corretamente
	assert.Error(t, err)
	assert.Nil(t, documents)
}

// Teste do método UpdateBlockDocument
func TestUpdateBlockDocument(t *testing.T) {
	// Criando uma instância do mock
	mockRepo := new(MockRepository)

	// Número do documento a ser bloqueado
	documentNumber := "123.456.789-00"

	// Teste de sucesso: configurando o mock para retornar nil (sem erro)
	mockRepo.On("UpdateBlockDocument", documentNumber).Return(nil).Once()

	// Chamando o método UpdateBlockDocument
	err := mockRepo.UpdateBlockDocument(documentNumber)

	// Verificando se o erro foi nil (sem erro)
	assert.NoError(t, err)

	// Teste de erro: configurando o mock para retornar um erro
	mockRepo.On("UpdateBlockDocument", documentNumber).Return(errors.New("erro ao bloquear documento")).Once()

	// Chamando o método UpdateBlockDocument novamente e verificando o erro
	err = mockRepo.UpdateBlockDocument(documentNumber)
	assert.Error(t, err)
	assert.Equal(t, "erro ao bloquear documento", err.Error())
}
