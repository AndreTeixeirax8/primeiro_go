package entity

import (
	"github.com/primeiro/internal/autenticacao/domain/validator"
	"github.com/rs/xid"
)

type Unidade struct {
	ID       string `json:"id"`        // ID da unidade
	Nome     string `validate:"required, min=3, max=100" json:"nome"`      // Nome da unidade
	Cnpj     string `json:"cnpj"`      // CNPJ da unidade
	Email    string `validate:"email" json:"email"`     // Email da unidade
	QtdSilos int    `validate:"min=0" json:"qtd_silos"` // Quantidade de silos da unidade
}

func NweUnidade( nome, cnpj, email string, qtdSilos int) (*Unidade,error) {
	unidade := &Unidade{
		ID:       xid.New().String(),
		Nome:     nome,
		Cnpj:     cnpj,
		Email:    email,
		QtdSilos: qtdSilos,
	}
	err := validator.ValidateStruct(unidade)
	if err != nil {	
		return nil,err
	}
	return unidade,nil
}