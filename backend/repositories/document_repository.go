package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"safeCpf/config"
	"safeCpf/models"
)

// DocumentRepository
type DocumentRepository interface {
	CreateDocument(cpf models.DocumentModel) error
	GetDocuments() ([]models.DocumentModel, error)
	UpdateBlockDocument(numberDocument string) error
}

// DocumentRepositoryImpl implementa a interface DocumentRepository
type DocumentRepositoryImpl struct {
	collection *mongo.Collection
}

// NewDocumentRepository
func NewDocumentRepository(collection *mongo.Collection) *DocumentRepositoryImpl {
	return &DocumentRepositoryImpl{collection: collection}
}

// CreateDocument insere um novo Document no banco de dados
func (r *DocumentRepositoryImpl) CreateDocument(cpf models.DocumentModel) error {
	_, err := r.collection.InsertOne(context.TODO(), cpf)
	if err != nil {
		return err
	}
	return nil
}

// GetDocuments consulta e retorna uma lista de registros de documents
func (r *DocumentRepositoryImpl) GetDocuments() ([]models.DocumentModel, error) {
	collection := config.GetCollection("documents")
	cursor, err := collection.Find(nil, bson.M{"block": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}

	var documents []models.DocumentModel
	if err := cursor.All(nil, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}

// UpdateBlockDocument Atualiza um document pelo número
func (r *DocumentRepositoryImpl) UpdateBlockDocument(numberDocument string) error {
	collection := config.GetCollection("documents")
	_, err := collection.UpdateOne(nil, bson.M{"numero": numberDocument}, bson.M{"$set": bson.M{"block": true}})
	if err != nil {
		return err
	}
	return nil
}
