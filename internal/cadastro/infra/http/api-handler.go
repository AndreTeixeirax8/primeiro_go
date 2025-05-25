package apiHandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/primeiro/internal/cadastro/infra/http/handler"
	"github.com/primeiro/pkg/middleware"
	"github.com/primeiro/pkg/wrapper"
)

type ApiHttpHandler struct {
	r *chi.Mux	
}

func NewApiHttpHandler(r *chi.Mux) *ApiHttpHandler {
	
	return &ApiHttpHandler{r: r}
}

func (s *ApiHttpHandler)RunAutenticacaoApi() {
	s.r.Route("/autenticacao", func(r chi.Router) {
		
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			w.Write([]byte("Autenticacao API"))
		})
		
	})


}

func (s *ApiHttpHandler)RunCadastroApi() {

	unidadeHandler := handler.NewUnidadeHandler()

	s.r.Route("/cadastro", func(r chi.Router) {
		
		r.Use(middleware.ApiKey)
		r.Route("/unidade", func(r chi.Router) {
			r.Post("/", wrapper.HandleError(unidadeHandler.CreateUnidade) )
			r.Get("/", wrapper.HandleError(unidadeHandler.ListUnidades) )
			r.Get("/{id}", wrapper.HandleError(unidadeHandler.GetUnidadeById) )
			})
		
	
	})


}