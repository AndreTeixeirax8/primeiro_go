package usecase

import "github.com/primeiro/internal/modules/autenticacao/domain/entity"

type GetUnidadeByIdInputDTO struct {
	ID string `json:"id"` // ID da unidade
}

type GetUnidadeByIdOutputDTO struct {
	ID       string `json:"id"`        // ID da unidade
	Nome     string `json:"nome"`      // Nome da unidade
	Cnpj     string `json:"cnpj"`      // CNPJ da unidade
	Email    string `json:"email"`     // Email da unidade
	QtdSilos int    `json:"qtd_silos"` // Quantidade de silos da unidade
}
type GetUnidadeByIdUsecase struct {
}

func NewGetUnidadeByIdUsecase() *GetUnidadeByIdUsecase {
	return &GetUnidadeByIdUsecase{}
}

func (uc *GetUnidadeByIdUsecase) Execute(input *GetUnidadeByIdInputDTO) (*GetUnidadeByIdOutputDTO, error) {

	unidadeMock := entity.Unidade{
		ID:       "1",
		Nome:     "Unidade 1",
		Cnpj:     "12345678000195",
		Email:    "joao@teste.com.br",
		QtdSilos: 10,
	}

	return &GetUnidadeByIdOutputDTO{
		ID:       unidadeMock.ID,
		Nome:     unidadeMock.Nome,
		Cnpj:     unidadeMock.Cnpj,
		Email:    unidadeMock.Email,
		QtdSilos: unidadeMock.QtdSilos,
	}, nil
}
