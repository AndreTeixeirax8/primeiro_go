package valueobject

import (
	"database/sql/driver"
	"errors"

	validator "github.com/Nhanderu/brdoc"
)

type CNPJ struct {
	valor string
}

func NewCNPJ(value string) (*CNPJ, error) {

	if !validator.IsCNPJ(value) {
		return nil, errors.New("CNPJ invalido")
	}

	return &CNPJ{valor: value}, nil
}

func (c *CNPJ) String() string {
	return c.valor
}

func (c *CNPJ) Scan(value interface{}) error {
	if value == nil {
		return nil

	}
	str, ok := value.(string)
	if !ok {
		return errors.New("CNPJ INVALIDO")
	}

	if str == "" {
		return nil
	}

	cnpj, err := NewCNPJ(str)
	if err != nil {
		return err
	}
	c.valor = cnpj.String()
	return nil
}

func (c *CNPJ) Value() (driver.Value, error) {
	return c.String(), nil
}
