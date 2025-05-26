package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	infrabot "github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/interface/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/start"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/stockinfo"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/fundinfo"
	"github.com/Rodrigoos/stock-bot-telegram/pkg/scraper"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN não definido")
	}

	bot := infrabot.NewTelegramBot(token)

	startUC := start.NewStartUseCase()

	scraper := scraper.NewStatusInvestScraper()
	stockUC := stockinfo.NewStockInfoUseCase(scraper)
	fundUC := fundinfo.NewFundInfoUseCase(scraper)

	handler := telegram.NewHandler(bot.API, startUC, stockUC, fundUC)

	log.Println("Bot iniciado...")
	handler.HandleUpdates()
}
