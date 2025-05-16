package main

import (
	"log"
	"os"

	infrabot "github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/interface/telegram"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/start"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/stockinfo"
	"github.com/Rodrigoos/stock-bot-telegram/pkg/scraper"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN não definido")
	}

	bot := infrabot.NewTelegramBot(token)

	startUC := start.NewStartUseCase()

	scraper := scraper.NewStatusInvestScraper()
	stockUC := stockinfo.NewStockInfoUseCase(scraper)

	handler := telegram.NewHandler(bot.API, startUC, stockUC)

	log.Println("Bot iniciado...")
	handler.HandleUpdates()
}
