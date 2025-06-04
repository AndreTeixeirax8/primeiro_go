package entity

import (
	"github.com/primeiro/internal/modules/cadastro/domain/validator"
	"github.com/primeiro/internal/modules/cadastro/domain/valueobject"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type Unidade struct {
	gorm.Model                   // Embedding gorm.Model for ID, CreatedAt, UpdatedAt, DeletedAt
	ID         string            `json:"id"` // ID da unidade
	Nome       string            `json:"nome" validate:"required,min=3,max=100"`
	Cnpj       *valueobject.CNPJ `json:"cnpj"`                       // CNPJ da unidade
	Email      string            `json:"email" validate:"email"`     // Email da unidade
	QtdSilos   int               `json:"qtd_silos" validate:"min=0"` // Quantidade de silos
}

func (Unidade) TableName() string {
	return "unidade"
}

func NewUnidade(nome, cnpj, email string, qtdSilos int) (*Unidade, error) {

	cnpjVo, err := valueobject.NewCNPJ(cnpj)
	if err != nil {
		return nil, err
	}

	unidade := &Unidade{
		ID:       xid.New().String(),
		Nome:     nome,
		Cnpj:     cnpjVo,
		Email:    email,
		QtdSilos: qtdSilos,
	}
	err = validator.ValidateStruct(unidade)
	if err != nil {
		return nil, err
	}
	return unidade, nil
}
