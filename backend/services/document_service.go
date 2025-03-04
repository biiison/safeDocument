package services

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
	"safeCpf/dto"
	"safeCpf/models"
	"safeCpf/repositories"
	"strconv"
	"strings"
	"time"
)

// DocumentService
type DocumentService interface {
	CreateDocument(document dto.CreateDocumentDto) error
	GetDocuments() ([]models.DocumentModel, error)
	UpdateBlockDocument(numberDocument string) error
}

// DocumentServiceImpl implementa a interface DocumentService
type DocumentServiceImpl struct {
	CpfRepository repositories.DocumentRepository // Aqui, a interface está sendo usada corretamente
}

// NewDocumentService
func NewDocumentService(cpfRepository repositories.DocumentRepository) *DocumentServiceImpl {
	return &DocumentServiceImpl{CpfRepository: cpfRepository}
}

// CreateDocument valida e cria um novo Document
func (s *DocumentServiceImpl) CreateDocument(documentReq dto.CreateDocumentDto) error {
	document := models.DocumentModel{
		Numero:       documentReq.Numero,
		Tipo:         documentReq.Tipo,
		Valido:       true,
		Block:        false,
		DataInclusao: time.Now(),
	}

	document.Valido = ValidaDocument(document.Numero, document.Tipo)
	document.DataInclusao = time.Now()

	if !document.Valido {
		return fmt.Errorf("%s invalido!", strings.ToUpper(document.Tipo))
	}

	err := s.CpfRepository.CreateDocument(document)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return &CustomError{Message: fmt.Sprintf("%s já cadastrado anteriormente!", strings.ToUpper(document.Tipo)), StatusCode: 409}
		}
		return err
	}
	return nil
}

func (s *DocumentServiceImpl) GetDocuments() ([]models.DocumentModel, error) {
	return s.CpfRepository.GetDocuments()
}

func (s *DocumentServiceImpl) UpdateBlockDocument(numberDocument string) error {
	return s.CpfRepository.UpdateBlockDocument(numberDocument)
}

type CustomError struct {
	Message    string
	StatusCode int
}

func (e *CustomError) Error() string {
	return e.Message
}
func ValidaDocument(numero, tipo string) bool {
	if tipo == "cpf" {
		re := regexp.MustCompile(`\D`)
		cpf := re.ReplaceAllString(numero, "")

		if len(cpf) != 11 {
			return false
		}

		calculaDigito := func(tamanho int) int {
			soma := 0
			for i := 0; i < tamanho; i++ {
				num, _ := strconv.Atoi(string(cpf[i]))
				soma += num * (tamanho + 1 - i)
			}
			resto := (soma * 10) % 11
			if resto == 10 {
				resto = 0
			}
			return resto
		}

		return calculaDigito(9) == int(cpf[9]-'0') && calculaDigito(10) == int(cpf[10]-'0')
	}
	if tipo == "cnpj" {
		cnpj := strings.Join(strings.Fields(numero), "")
		if len(cnpj) != 14 {
			return false
		}

		if cnpj == "00000000000000" || cnpj == "11111111111111" || cnpj == "22222222222222" || cnpj == "33333333333333" ||
			cnpj == "44444444444444" || cnpj == "55555555555555" || cnpj == "66666666666666" || cnpj == "77777777777777" ||
			cnpj == "88888888888888" || cnpj == "99999999999999" {
			return false
		}

		calcDigito := func(cnpj string, pesos []int) int {
			soma := 0
			for i, peso := range pesos {
				soma += int(cnpj[i]-'0') * peso
			}
			resultado := soma % 11
			if resultado < 2 {
				return 0
			}
			return 11 - resultado
		}

		pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

		d1 := calcDigito(cnpj, pesos1)
		d2 := calcDigito(cnpj, pesos2)

		return int(cnpj[12]-'0') == d1 && int(cnpj[13]-'0') == d2
	}
	return false
}
