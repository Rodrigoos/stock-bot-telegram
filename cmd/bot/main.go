package main

import (
	"log"
	"os"

	"github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/db"
	infrabot "github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/interface/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/fundinfo"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/start"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/stockinfo"
	"github.com/Rodrigoos/stock-bot-telegram/pkg/scraper"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN não definido")
	}

	db.Connect()

	db.DB.AutoMigrate(&models.Portfolio{}, &models.Asset{})

	bot := infrabot.NewTelegramBot(token)

	startUC := start.NewStartUseCase()

	portfolio := models.Portfolio{
		Name: "Minha Carteira FII",
		Assets: []models.Asset{
			{Ticker: "HGLG11", Quantity: 10, Price: 158.0},
			{Ticker: "KNRI11", Quantity: 5, Price: 125.5},
		},
	}

	result := db.DB.Create(&portfolio)
	if result.Error != nil {
		log.Println("Erro ao criar portfólio:", result.Error)
	}

	scraper := scraper.NewStatusInvestScraper()
	stockUC := stockinfo.NewStockInfoUseCase(scraper)
	fundUC := fundinfo.NewFundInfoUseCase(scraper)

	handler := telegram.NewHandler(bot.API, startUC, stockUC, fundUC)

	log.Println("Bot iniciado...")
	handler.HandleUpdates()
}
