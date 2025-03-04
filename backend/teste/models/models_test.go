package models

import (
	"github.com/stretchr/testify/assert"
	"safeCpf/models"
	"testing"
	"time"
)

func TestDocumentModel(t *testing.T) {
	t.Run("Test DocumentModel Initialization", func(t *testing.T) {
		dataInclusao := time.Now()
		document := models.DocumentModel{
			Numero:       "123.456.789-00",
			Tipo:         "cpf",
			Valido:       true,
			Block:        false,
			DataInclusao: dataInclusao,
		}

		assert.Equal(t, "123.456.789-00", document.Numero)
		assert.Equal(t, "cpf", document.Tipo)
		assert.Equal(t, true, document.Valido)
		assert.Equal(t, false, document.Block)
		assert.Equal(t, dataInclusao.Format(time.RFC3339), document.DataInclusao.Format(time.RFC3339))
	})

	t.Run("Test DataInclusao", func(t *testing.T) {
		fixedDate := time.Date(2025, time.March, 4, 12, 30, 0, 0, time.UTC)

		document := models.DocumentModel{
			Numero:       "987.654.321-00",
			Tipo:         "rg",
			Valido:       false,
			Block:        true,
			DataInclusao: fixedDate,
		}

		assert.Equal(t, fixedDate.Format(time.RFC3339), document.DataInclusao.Format(time.RFC3339))
	})

	t.Run("Test Default Block Value", func(t *testing.T) {
		document := models.DocumentModel{
			Numero:       "111.222.333-44",
			Tipo:         "cpf",
			Valido:       true,
			DataInclusao: time.Now(),
		}

		assert.Equal(t, false, document.Block)
	})
}
