package main

import (
	"log"

	"github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/db"
	"github.com/Rodrigoos/stock-bot-telegram/internal/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	db.Connect()

	err = db.DB.AutoMigrate(&models.Portfolio{}, &models.Asset{})
	if err != nil {
		log.Fatal("Erro ao migrar: ", err)
	}

	log.Println("Migrations realizadas com sucesso!")
}
