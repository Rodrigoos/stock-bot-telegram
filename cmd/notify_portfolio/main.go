package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/db"
	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
	"github.com/Rodrigoos/stock-bot-telegram/internal/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não carregado")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatIDStr == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN ou TELEGRAM_CHAT_ID não definidos")
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("CHAT_ID inválido: %v", err)
	}

	db := db.Connect()

	if len(os.Args) < 2 {
		log.Fatal("Uso: go run ./cmd/notifyportfolio <portfolio_id>")
	}

	portfolioID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("ID inválido:", err)
	}

	var portfolio models.Portfolio
	err = db.Preload("Assets").First(&portfolio, portfolioID).Error
	if err != nil {
		log.Fatalf("Erro ao buscar carteira: %v", err)
	}

	message := utils.FormatPortfolioMessage(portfolio)

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("Erro ao criar bot:", err)
	}

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "Markdown"
	bot.Send(msg)

	log.Println("Mensagem enviada com sucesso.")
}
