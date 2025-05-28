package domain

import (
	"github.com/primeiro/internal/modules/cadastro/domain/entity"
	pkg "github.com/primeiro/pkg/repository"
)

type UnidadeRepository interface {
	pkg.RepositoryBaseInterface[entity.Unidade]
}
