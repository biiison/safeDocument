package models

import "time"

type DocumentModel struct {
	Numero       string    `json:"numero"`
	Tipo         string    `json:"tipo"`
	Valido       bool      `json:"valido"`
	Block        bool      `json:"block"`
	DataInclusao time.Time `json:"data_inclusao"`
}
