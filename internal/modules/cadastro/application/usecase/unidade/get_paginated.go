package usecase

import (
	domain "github.com/primeiro/internal/modules/cadastro/domain/repository"
	"github.com/primeiro/pkg/pagination"
)

type ContatosDTO struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
type GetUnidadePaginatedOutputDTO struct {
	ID       string        `json:"id"`        // ID da unidade
	Nome     string        `json:"nome"`      // Nome da unidade
	Cnpj     string        `json:"cnpj"`      // CNPJ da unidade
	Email    string        `json:"email"`     // Email da unidade
	QtdSilos int           `json:"qtd_silos"` // Quantidade de silos da unidade
	Contatos []ContatosDTO `json:"contatos"`
}
type GetUnidadePaginatedUsecase struct {
	repo domain.UnidadeRepository
}

func NewGetUnidadePaginatedUsecase(repo domain.UnidadeRepository) *GetUnidadePaginatedUsecase {
	return &GetUnidadePaginatedUsecase{repo: repo}
}

func (uc *GetUnidadePaginatedUsecase) Execute(input *pagination.PaginationQuery) (*pagination.PaginationResponse[GetUnidadePaginatedOutputDTO], error) {

	ret, err := uc.repo.GetPaginated(input, "Contatos")
	if err != nil {
		return nil, err
	}

	rows := make([]GetUnidadePaginatedOutputDTO, len(ret.Rows))
	for i, unidade := range ret.Rows {

		var contatos []ContatosDTO
		for _, contato := range unidade.Contatos {
			contatos = append(contatos, ContatosDTO{
				ID:    contato.ID,
				Nome:  contato.Nome,
				Email: contato.Email,
			})
		}

		rows[i] = GetUnidadePaginatedOutputDTO{
			ID:       unidade.ID,
			Nome:     unidade.Nome,
			Cnpj:     unidade.Cnpj.String(),
			Email:    unidade.Email,
			QtdSilos: unidade.QtdSilos,
			Contatos: contatos,
		}
	}

	return &pagination.PaginationResponse[GetUnidadePaginatedOutputDTO]{
		Rows: rows,
		Meta: ret.Meta,
	}, nil

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
