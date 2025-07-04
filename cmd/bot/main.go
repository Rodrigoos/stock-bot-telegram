package main

import (
	"log"
	"os"

	"github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/db"
	infrabot "github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/interface/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase"
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

	status_scraper := scraper.NewStatusInvestScraper()
	binance_scraper := scraper.NewBinanceScraper()

	startUC := usecase.NewStartUseCase()
	stockUC := usecase.NewStockInfoUseCase(status_scraper)
	fundUC := usecase.NewFundInfoUseCase(status_scraper)
	criptoUC := usecase.NewCriptoInfoUseCase(binance_scraper)
	portfolioService := usecase.NewPortfolioService(db.Connect())

	handler := telegram.NewHandler(bot.API,
		startUC,
		stockUC,
		fundUC,
		criptoUC,
		portfolioService,
	)

	log.Println("Bot iniciado...")
	handler.HandleUpdates()
}
