package usecase

import (
	domain "github.com/primeiro/internal/modules/cadastro/domain/repository"
)

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
	repo domain.UnidadeRepository
}

func NewGetUnidadeByIdUsecase(repo domain.UnidadeRepository) *GetUnidadeByIdUsecase {
	return &GetUnidadeByIdUsecase{repo: repo}
}

func (uc *GetUnidadeByIdUsecase) Execute(input *GetUnidadeByIdInputDTO) (*GetUnidadeByIdOutputDTO, error) {

	unidade, err := uc.repo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetUnidadeByIdOutputDTO{
		ID:       unidade.ID,
		Nome:     unidade.Nome,
		Cnpj:     unidade.Cnpj.String(),
		Email:    unidade.Email,
		QtdSilos: unidade.QtdSilos,
	}, nil
}
