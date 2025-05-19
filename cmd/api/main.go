package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	apiHandler "github.com/primeiro/internal/autenticacao/infra/http"
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
	if servicoAutenticacaoAtivo {
		apiHandler.RunAutenticacaoApi()
	} else {
		logger.Println("Serviço de autenticação inativo")
	}

}
