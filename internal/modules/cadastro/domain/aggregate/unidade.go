package aggregate

import (
	"errors"

	"github.com/primeiro/internal/modules/cadastro/domain/entity"
)

type Unidade struct {
	entity.Unidade
	contatos []entity.Contato
}

func NewUnidade(nome, cnpj, email string) (*Unidade, error) {
	unidade, err := entity.NewUnidade(nome, cnpj, email, 0)
	if err != nil {
		return nil, err
	}

	unidadeAggregate := &Unidade{Unidade: *unidade}

	return unidadeAggregate, nil
}

func (u *Unidade) AddContato(nome, email string) error {

	contato, err := entity.NewContato(u.ID, nome, email)
	if err != nil {
		return err
	}

	if u.contatos == nil {
		u.contatos = make([]entity.Contato, 0)
	}

	u.contatos = append(u.contatos, *contato)
	return nil
}

func (u *Unidade) GetContatos() []entity.Contato {
	return u.contatos
}

func (u *Unidade) Validate() error {
	if len(u.contatos) == 0 {
		return errors.New("Unidade sem contatos")
	}
	return nil
}
