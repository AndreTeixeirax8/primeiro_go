package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	apiHandlerCadastro "github.com/primeiro/internal/cadastro/infra/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	logger := log.New(os.Stdout, "api: ", log.LstdFlags)
	_, isDocker := os.LookupEnv("DOCKER")
	if !isDocker {
		if err := godotenv.Load(); err != nil {
			logger.Println("Error loading .env file")
			panic(err)
		}
	}




	servicoAutenticacaoAtivo , err :=  strconv.ParseBool(os.Getenv("SERVICO_AUTENTICACAO"))
	if err != nil {
		servicoAutenticacaoAtivo =true
	}
	
	servicoCadastroAtivo , err :=  strconv.ParseBool(os.Getenv("SERVICO_CADASTRO"))
	if err != nil {
		servicoCadastroAtivo =true
	}
	

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
	     ExposedHeaders:   []string{"Link"}, // Add this line to expose the Link header
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	

	apiHandler := apiHandlerCadastro.NewApiHttpHandler(r)

	if servicoAutenticacaoAtivo {
		apiHandler.RunAutenticacaoApi()
	} else {
		logger.Println("Serviço de autenticação inativo")
	}

	if servicoCadastroAtivo {
		apiHandler.RunCadastroApi()
	} else {
		logger.Println("Serviço de cadastro inativo")
	}


	logger.Println("Iniciando servidor na porta " + os.Getenv("PORT"))
	// Inicia o servidor HTTP
	err= http.ListenAndServe(":" + os.Getenv("PORT"), r)
	if err != nil {
		logger.Println("Erro ao iniciar servidor: ", err)
		panic(err)
	}
}
