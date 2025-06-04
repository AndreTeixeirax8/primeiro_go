package usecase

import (
	domain "github.com/primeiro/internal/modules/cadastro/domain/repository"
)

type ListUnidadeOutputDTO struct {
	ID       string `json:"id"`        // ID da unidade
	Nome     string `json:"nome"`      // Nome da unidade
	Cnpj     string `json:"cnpj"`      // CNPJ da unidade
	Email    string `json:"email"`     // Email da unidade
	QtdSilos int    `json:"qtd_silos"` // Quantidade de silos da unidade
}
type ListUnidadeUsecase struct {
	repo domain.UnidadeRepository
}

func NewListUnidadesUsecase(repo domain.UnidadeRepository) *ListUnidadeUsecase {
	return &ListUnidadeUsecase{repo: repo}
}

func (uc *ListUnidadeUsecase) Execute() (*[]ListUnidadeOutputDTO, error) {

	unidades, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	rows := make([]ListUnidadeOutputDTO, len(unidades))
	for i, unidade := range unidades {
		rows[i] = ListUnidadeOutputDTO{
			ID:       unidade.ID,
			Nome:     unidade.Nome,
			Cnpj:     unidade.Cnpj.String(),
			Email:    unidade.Email,
			QtdSilos: unidade.QtdSilos,
		}
	}

	return &rows, nil
}
