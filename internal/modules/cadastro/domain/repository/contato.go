package domain

import (
	"github.com/primeiro/internal/modules/cadastro/domain/entity"
	pkg "github.com/primeiro/pkg/repository"
)

type ContatoRepository interface {
	pkg.RepositoryBaseInterface[entity.Contato]
}
