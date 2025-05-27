package usecase

import "github.com/primeiro/internal/modules/autenticacao/domain/entity"

type CreateUnidadeInputDTO struct {
	Nome     string `json:"nome"`      // Nome da unidade
	Cnpj     string `json:"cnpj"`      // CNPJ da unidade
	Email    string `json:"email"`     // Email da unidade
	QtdSilos int    `json:"qtd_silos"` // Quantidade de silos da unidade
}

type CreateUnidadeOutputDTO struct {
	ID string `json:"id"` // ID da unidade
}
type CreateUnidadeUseCase struct {
}

func NewCreateUnidadeUseCase() *CreateUnidadeUseCase {
	return &CreateUnidadeUseCase{}
}

func (uc *CreateUnidadeUseCase) Execute(input *CreateUnidadeInputDTO) (*CreateUnidadeOutputDTO, error) {
	unidade, err := entity.NewUnidade(input.Nome, input.Cnpj, input.Email, input.QtdSilos)
	if err != nil {
		return nil, err
	}
	return &CreateUnidadeOutputDTO{
		ID: unidade.ID}, nil
}
