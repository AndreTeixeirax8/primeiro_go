package handler

import (
	"encoding/json"
	"net/http"

	usecase "github.com/primeiro/internal/autenticacao/application/usecase/unidade"
)

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

func (h *UnidadeHandler) CreateUnidade(w http.ResponseWriter, r *http.Request)(interface{},int,error) {
	var input usecase.CreateUnidadeInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	output, err := h.createUnidadeUsecase.Execute(&input)
	
	return output, http.StatusCreated, err
}