package usecase

import (
	"github.com/primeiro/internal/modules/cadastro/domain/aggregate"
	domain "github.com/primeiro/internal/modules/cadastro/domain/repository"
)

type CreateUnidadeAggregateInputDTO struct {
	Nome     string `json:"nome"`  // Nome da unidade
	Cnpj     string `json:"cnpj"`  // CNPJ da unidade
	Email    string `json:"email"` // Email da unidade
	Contatos []struct {
		Nome  string `json:"nome"`
		Email string `json:"email"`
	} `json:"contatos"`
}

type CreateUnidadeAggregateOutputDTO struct {
	ID string `json:"id"` // ID da unidade
}
type CreateUnidadeAggregateUseCase struct {
	repo        domain.UnidadeRepository
	repoContato domain.ContatoRepository
}

func NewCreateUnidadeAggregateUseCase(
	repo domain.UnidadeRepository,
	repoContato domain.ContatoRepository,
) *CreateUnidadeAggregateUseCase {
	return &CreateUnidadeAggregateUseCase{repo: repo, repoContato: repoContato}
}

func (uc *CreateUnidadeAggregateUseCase) Execute(input *CreateUnidadeAggregateInputDTO) (*CreateUnidadeAggregateOutputDTO, error) {
	unidade, err := aggregate.NewUnidade(input.Nome, input.Cnpj, input.Email)
	if err != nil {
		return nil, err
	}

	for _, contato := range input.Contatos {
		err = unidade.AddContato(contato.Nome, contato.Email)
		if err != nil {
			return nil, err
		}
	}

	err = unidade.Validate()
	if err != nil {
		return nil, err
	}

	tx := uc.repo.BeginTx()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = uc.repo.CreateTx(unidade.GetUnidade(), tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 🔥 Adiciona isso pra salvar os contatos:
	for _, contato := range unidade.GetContatos() {
		err = uc.repoContato.CreateTx(&contato, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &CreateUnidadeAggregateOutputDTO{ID: unidade.ID}, nil
}
