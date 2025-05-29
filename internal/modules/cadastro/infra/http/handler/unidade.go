package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/primeiro/internal/infra/database"
	repository "github.com/primeiro/internal/infra/database/repository/cadastro"
	usecase "github.com/primeiro/internal/modules/cadastro/application/usecase/unidade"
	"github.com/primeiro/pkg/pagination"
)

type UnidadeHandler struct {
	createUnidadeUsecase       *usecase.CreateUnidadeUseCase
	getUnidadeByIdUsecase      *usecase.GetUnidadeByIdUsecase
	listUnidadeUsecase         *usecase.ListUnidadeUsecase
	getUnidadePaginatedUsecase *usecase.GetUnidadePaginatedUsecase
}

func NewUnidadeHandler() *UnidadeHandler {

	unidadeRepo := repository.NewUnidadeRepository(database.DB)

	createUnidadeUsecase := usecase.NewCreateUnidadeUseCase(unidadeRepo)
	getUnidadeByIdUsecase := usecase.NewGetUnidadeByIdUsecase(unidadeRepo)
	listUnidadeUsecase := usecase.NewListUnidadesUsecase(unidadeRepo)
	getUnidadePaginatedUsecase := usecase.NewGetUnidadePaginatedUsecase(unidadeRepo)

	return &UnidadeHandler{
		createUnidadeUsecase:       createUnidadeUsecase,
		getUnidadeByIdUsecase:      getUnidadeByIdUsecase,
		listUnidadeUsecase:         listUnidadeUsecase,
		getUnidadePaginatedUsecase: getUnidadePaginatedUsecase,
	}
}

func (h *UnidadeHandler) CreateUnidade(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input usecase.CreateUnidadeInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	output, err := h.createUnidadeUsecase.Execute(&input)

	return output, http.StatusCreated, err
}

func (h *UnidadeHandler) GetUnidadeById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := chi.URLParam(r, "id")
	input := usecase.GetUnidadeByIdInputDTO{ID: id}
	output, err := h.getUnidadeByIdUsecase.Execute(&input)
	return output, http.StatusOK, err
}

func (h *UnidadeHandler) ListUnidades(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	output, err := h.listUnidadeUsecase.Execute()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, http.StatusOK, nil
}

func (h *UnidadeHandler) GetUnidadeByPaginated(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	query := r.URL.Query()

	paginationRequest := pagination.GeneratePaginationRequest(query)

	output, err := h.getUnidadePaginatedUsecase.Execute(*paginationRequest)

	return output, http.StatusOK, err

}
