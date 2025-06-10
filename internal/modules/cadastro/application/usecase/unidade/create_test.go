package usecase_test

import (
	"testing"

	usecase "github.com/primeiro/internal/modules/cadastro/application/usecase/unidade"
	"github.com/primeiro/internal/modules/cadastro/domain/entity"
	domain "github.com/primeiro/internal/modules/cadastro/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUnidadeRepository struct {
	mock.Mock
	domain.UnidadeRepository
}

func (m *MockUnidadeRepository) Create(unidade *entity.Unidade) (*entity.Unidade, error) {
	args := m.Called(unidade)
	return args.Get(0).(*entity.Unidade), args.Error(1)
}

func TestCreateUnidadeUsecase_Success(t *testing.T) {
	repo := new(MockUnidadeRepository)
	mockUnidade := &entity.Unidade{ID: "mock-id"}
	repo.On("Create", mock.Anything).Return(mockUnidade, nil)

	uc := usecase.NewCreateUnidadeUseCase(repo)
	input := &usecase.CreateUnidadeInputDTO{
		Nome:     "Unidade Teste",
		Cnpj:     "12.345.678/0001-95",
		Email:    "unidade@teste.com.br",
		QtdSilos: 10}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	repo.AssertCalled(t, "Create", mock.Anything)
	repo.AssertNumberOfCalls(t, "Create", 1)
}
