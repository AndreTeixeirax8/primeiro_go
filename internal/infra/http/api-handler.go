package apiHandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	handler "github.com/primeiro/internal/infra/http/cadastro"
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

		r.Route("/unidade", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET unidade Autenticacao API"))
		})

		r.Post("/", wrapper.HandleError(unidadeHandler.CreateUnidade) )

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get id Unidade Autenticacao API"))
		})


		})
		
	})


}