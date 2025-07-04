package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

	message := formatPortfolioMessage(portfolio)

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("Erro ao criar bot:", err)
	}

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "Markdown"
	bot.Send(msg)

	log.Println("Mensagem enviada com sucesso.")
}

// func formatPortfolioMessage(p models.Portfolio) string {
// 	if len(p.Assets) == 0 {
// 		return fmt.Sprintf("*%s*\nNenhum ativo cadastrado.", p.Name)
// 	}

// 	var b strings.Builder
// 	b.WriteString(fmt.Sprintf("*%s*\n\n", p.Name))

// 	for _, asset := range p.Assets {
// 		total := float64(asset.Quantity) * asset.Price
// 		b.WriteString(fmt.Sprintf("• *%s*: %d × R$ %.2f = R$ %.2f\n", asset.Ticker, asset.Quantity, asset.Price, total))
// 	}

// 	return b.String()
// }

func formatPortfolioMessage(p models.Portfolio) string {
	if len(p.Assets) == 0 {
		return fmt.Sprintf("Carteira: %s\n\nNenhum ativo encontrado.", p.Name)
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("Carteira: %s\n\n", p.Name))
	b.WriteString("```\n")
	b.WriteString(fmt.Sprintf("%-7s %-6s %-12s %-12s\n", "Ticker", "Qtde", "Preço", "Total"))
	b.WriteString(strings.Repeat("-", 42) + "\n")

	var totalGeral float64

	for _, asset := range p.Assets {
		total := float64(asset.Quantity) * asset.Price
		b.WriteString(fmt.Sprintf(
			"%-7s %-6d %-12s %-12s\n",
			asset.Ticker,
			asset.Quantity,
			utils.FormatBRL(asset.Price),
			utils.FormatBRL(total),
		))
		totalGeral += total
	}

	b.WriteString(strings.Repeat("-", 42) + "\n")
	b.WriteString(fmt.Sprintf("Total geral: %s\n", utils.FormatBRL(totalGeral)))
	b.WriteString("```")

	return b.String()
}
