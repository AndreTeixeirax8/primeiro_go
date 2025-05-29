package usecase

import (
	domain "github.com/primeiro/internal/modules/cadastro/domain/repository"
	"github.com/primeiro/pkg/pagination"
)

type GetUnidadePaginatedInputDTO struct {
	ID string `json:"id"` // ID da unidade
}

type GetUnidadePaginatedOutputDTO struct {
	ID       string `json:"id"`        // ID da unidade
	Nome     string `json:"nome"`      // Nome da unidade
	Cnpj     string `json:"cnpj"`      // CNPJ da unidade
	Email    string `json:"email"`     // Email da unidade
	QtdSilos int    `json:"qtd_silos"` // Quantidade de silos da unidade
}
type GetUnidadePaginatedUsecase struct {
	repo domain.UnidadeRepository
}

func NewGetUnidadePaginatedUsecase(repo domain.UnidadeRepository) *GetUnidadePaginatedUsecase {
	return &GetUnidadePaginatedUsecase{repo: repo}
}

func (uc *GetUnidadePaginatedUsecase) Execute(input pagination.PaginationQuery) (*GetUnidadePaginatedOutputDTO, error) {
	/*
		unidade, err := uc.repo.GetPaginated(input.ID)
		if err != nil {
			return nil, err
		}

		return &GetUnidadePaginatedOutputDTO{
			ID:       unidade.ID,
			Nome:     unidade.Nome,
			Cnpj:     unidade.Cnpj,
			Email:    unidade.Email,
			QtdSilos: unidade.QtdSilos,
		}, nil*/
}
