package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL não configurada ")
	}

	config := &gorm.Config{}

	if os.Getenv("DATABASE_DEBUG") == "true" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")

}
