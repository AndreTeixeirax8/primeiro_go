package usecase

import "github.com/primeiro/internal/modules/autenticacao/domain/entity"

type ListUnidadeOutputDTO struct {
	ID       string `json:"id"`        // ID da unidade
	Nome     string `json:"nome"`      // Nome da unidade
	Cnpj     string `json:"cnpj"`      // CNPJ da unidade
	Email    string `json:"email"`     // Email da unidade
	QtdSilos int    `json:"qtd_silos"` // Quantidade de silos da unidade
}
type ListUnidadeUsecase struct {
}

func NewListUnidadesUsecase() *ListUnidadeUsecase {
	return &ListUnidadeUsecase{}
}

func (uc *ListUnidadeUsecase) Execute() (*[]ListUnidadeOutputDTO, error) {

	unidadeMock := []entity.Unidade{
		{
			ID:       "1",
			Nome:     "Unidade 1",
			Cnpj:     "12345678000195",
			Email:    "joao@teste.com.br",
			QtdSilos: 10,
		},
		{
			ID:       "2",
			Nome:     "Unidade 2",
			Cnpj:     "123456780001952",
			Email:    "joao2@teste.com.br",
			QtdSilos: 8,
		},
	}

	unidades := make([]ListUnidadeOutputDTO, len(unidadeMock))
	for i, unidade := range unidadeMock {
		unidades[i] = ListUnidadeOutputDTO{
			ID:       unidade.ID,
			Nome:     unidade.Nome,
			Cnpj:     unidade.Cnpj,
			Email:    unidade.Email,
			QtdSilos: unidade.QtdSilos,
		}
	}

	return &unidades, nil
}
