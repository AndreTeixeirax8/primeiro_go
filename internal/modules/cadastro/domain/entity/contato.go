package entity

import (
	"github.com/primeiro/internal/modules/autenticacao/domain/validator"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type Contato struct {
	gorm.Model          // Embedding gorm.Model for ID, CreatedAt, UpdatedAt, DeletedAt
	ID         string   `json:"id"` // ID da unidade
	UnidadeID  string   `json:"unidade_id"`
	Unidade    *Unidade ` json:"unidade"`
	Nome       string   `json:"nome" validate:"required,min=3,max=100"`
	Email      string   `json:"email" validate:"email"` // Email da unidade

}

func (Contato) TableName() string {
	return "contato"
}

func NewContato(unidadeId, nome, email string) (*Contato, error) {
	contato := &Contato{
		ID:        xid.New().String(),
		Nome:      nome,
		UnidadeID: unidadeId,
		Email:     email,
	}
	err := validator.ValidateStruct(contato)
	if err != nil {
		return nil, err
	}
	return contato, nil
}
