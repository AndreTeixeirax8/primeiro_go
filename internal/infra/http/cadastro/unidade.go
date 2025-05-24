package handler

import usecase "github.com/primeiro/internal/autenticacao/application/usecase/unidade"

type UnidadeHandler struct {
	createUnidadeUsecase *usecase.CreateUnidadeUseCase
	getUnidadeByIdUsecase    *usecase.GetUnidadeByIdUsecase
	listUnidadeUsecase *usecase.ListUnidadeUsecase

}

func NewUnidadeHandler() *UnidadeHandler {
	createUnidadeUsecase := usecase.NewCreateUnidadeUseCase()
	getUnidadeByIdUsecase  := usecase.NewGetUnidadeByIdUsecase()
	listUnidadeUsecase := usecase.NewListUnidadesUsecase()

	return &UnidadeHandler{
		createUnidadeUsecase: createUnidadeUsecase,
		getUnidadeByIdUsecase: getUnidadeByIdUsecase,
		listUnidadeUsecase: listUnidadeUsecase,
	}
}