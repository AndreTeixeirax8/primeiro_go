package repository

import (
	"github.com/primeiro/internal/modules/cadastro/domain/entity"
	pkg "github.com/primeiro/pkg/repository"
	"gorm.io/gorm"
)

type ContatoRepository struct {
	pkg.RepositoryBase[entity.Contato]
	//Db *gorm.DB
	SearchExpression string
}

func NewContatoRepository(db *gorm.DB) *ContatoRepository {
	return &ContatoRepository{
		RepositoryBase: pkg.RepositoryBase[entity.Contato]{Db: db},
	}
}
