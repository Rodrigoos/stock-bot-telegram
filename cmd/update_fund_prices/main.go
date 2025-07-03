package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/db"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/updateprices"
	"github.com/Rodrigoos/stock-bot-telegram/pkg/scraper"
)

func main() {
	godotenv.Load()

	scraper := scraper.NewStatusInvestScraper()
	service := updateprices.NewService(db.Connect(), scraper)

	// recebe ID da carteira via argumento
	if len(os.Args) < 2 {
		log.Fatal("Informe o ID da carteira: ex: go run ./cmd/updateprices 1")
	}
	id, _ := strconv.Atoi(os.Args[1])

	err := service.ExecuteFund(uint(id))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Preços atualizados com sucesso.")
}
