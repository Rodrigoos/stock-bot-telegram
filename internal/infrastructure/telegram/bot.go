package telegram

import (
	"fmt"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	API *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	const (
		maxRetries = 5
		timeout    = 10 * time.Second
	)

	httpClient := &http.Client{
		Timeout: timeout,
	}

	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		bot, err := tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, httpClient)
		if err == nil {
			log.Println("✅ Bot do Telegram conectado com sucesso")
			return &TelegramBot{API: bot}, nil
		}

		lastErr = err
		wait := time.Duration(attempt*2) * time.Second

		log.Printf(
			"⚠️ Tentativa %d/%d falhou ao conectar no Telegram: %v. Retry em %s...",
			attempt,
			maxRetries,
			err,
			wait,
		)

		time.Sleep(wait)
	}

	return nil, fmt.Errorf("falha ao conectar no Telegram após %d tentativas: %w", maxRetries, lastErr)
}
