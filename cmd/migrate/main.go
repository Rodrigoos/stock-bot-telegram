package main

import (
	"log"

	"stockbot/internal/infrastructure/database"
	"stockbot/internal/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	db := database.Connect()

	err = db.AutoMigrate(&models.Portfolio{}, &models.Asset{})
	if err != nil {
		log.Fatal("Erro ao migrar: ", err)
	}

	log.Println("Migrations realizadas com sucesso!")
}
