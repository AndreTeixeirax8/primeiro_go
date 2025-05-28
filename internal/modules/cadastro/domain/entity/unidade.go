package entity

import (
	"github.com/primeiro/internal/modules/autenticacao/domain/validator"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type Unidade struct {
	gorm.Model        // Embedding gorm.Model for ID, CreatedAt, UpdatedAt, DeletedAt
	ID         string `json:"id"` // ID da unidade
	Nome       string `json:"nome" validate:"required,min=3,max=100"`
	Cnpj       string `json:"cnpj"`                       // CNPJ da unidade
	Email      string `json:"email" validate:"email"`     // Email da unidade
	QtdSilos   int    `json:"qtd_silos" validate:"min=0"` // Quantidade de silos
}

func (Unidade) TableName() string {
	return "unidade"
}

func NewUnidade(nome, cnpj, email string, qtdSilos int) (*Unidade, error) {
	unidade := &Unidade{
		ID:       xid.New().String(),
		Nome:     nome,
		Cnpj:     cnpj,
		Email:    email,
		QtdSilos: qtdSilos,
	}
	err := validator.ValidateStruct(unidade)
	if err != nil {
		return nil, err
	}
	return unidade, nil
}
