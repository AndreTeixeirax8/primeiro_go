package entity_test

import (
	"testing"

	"github.com/primeiro/internal/modules/cadastro/domain/entity"
	"github.com/stretchr/testify/assert"
)

const (
	nome     = "Unidade A"
	cnpj     = "12.345.678/0001-95"
	email    = "unidade@teste.com"
	qtdSilos = 10
)

func TestNewUnidade_DadosValidos(t *testing.T) {

	unidade, err := entity.NewUnidade(nome, cnpj, email, qtdSilos)

	assert.NoError(t, err)
	assert.NotNil(t, unidade)
	assert.Equal(t, nome, unidade.Nome)
	assert.Equal(t, cnpj, unidade.Cnpj.String())
	assert.Equal(t, email, unidade.Email)
	assert.Equal(t, qtdSilos, unidade.QtdSilos)
}

func TestNewUnidade_CnpjInvalido(t *testing.T) {

	cnpjInvalido := "1235445"

	unidade, err := entity.NewUnidade(nome, cnpjInvalido, email, qtdSilos)

	assert.Error(t, err)
	assert.Nil(t, unidade)

}
func TestNewUnidade_EmailInvalido(t *testing.T) {

	emailInvalido := "email-invalido"

	unidade, err := entity.NewUnidade(nome, cnpj, emailInvalido, qtdSilos)

	assert.Error(t, err)
	assert.Nil(t, unidade)

}

func TestNewUnidade_SemNome(t *testing.T) {

	nomeInvalido := ""

	unidade, err := entity.NewUnidade(nome, cnpj, nomeInvalido, qtdSilos)

	assert.Error(t, err)
	assert.Nil(t, unidade)

}

func TestNewUnidade_NomeGrande(t *testing.T) {

	nomeInvalido := "12345678910123456789101234567891012345678910123456789101234567891012345678910123456789101234567891012345678910"

	unidade, err := entity.NewUnidade(nome, cnpj, nomeInvalido, qtdSilos)

	assert.Error(t, err)
	assert.Nil(t, unidade)

}
