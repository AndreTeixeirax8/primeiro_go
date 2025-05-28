package domain

import (
	"github.com/primeiro/internal/modules/cadastro/domain/entity"
)

type UnidadeRepository interface {
	Create(unidade *entity.Unidade) (*entity.Unidade, error)
	GetByID(id string) (*entity.Unidade, error)
	GetAll() ([]entity.Unidade, error)
	Update(unidade *entity.Unidade) (*entity.Unidade, error)
	Delete(id string) error
}
