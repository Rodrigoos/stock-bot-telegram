package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/start"
	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase/stockinfo"
)

type Handler struct {
	Bot              *tgbotapi.BotAPI
	StartUseCase     *start.StartUseCase
	StockInfoUseCase *stockinfo.StockInfoUseCase
}

func NewHandler(bot *tgbotapi.BotAPI, startUC *start.StartUseCase, stockUC *stockinfo.StockInfoUseCase) *Handler {
	return &Handler{	
		Bot:              bot,
		StartUseCase:     startUC,
		StockInfoUseCase: stockUC,
	}
}

func (h *Handler) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		text := strings.TrimSpace(update.Message.Text)

		switch {
		case text == "/start":
			msg := h.StartUseCase.Execute()
			h.sendMessage(update.Message.Chat.ID, msg)

		case strings.HasPrefix(text, "/stock"):
			args := strings.Fields(text)
			if len(args) < 2 {
				h.sendMessage(update.Message.Chat.ID, "Use o comando assim: /stock PETR4")
				continue
			}

			ticker := strings.ToUpper(args[1])

			info, err := h.StockInfoUseCase.Execute(ticker)
			if err != nil {
				h.sendMessage(update.Message.Chat.ID, "Erro ao buscar informação: "+err.Error())
				continue
			}

			h.sendMessage(update.Message.Chat.ID, info)

		default:
			h.sendMessage(update.Message.Chat.ID, "Comando não reconhecido.")
		}
	}
}

func (h *Handler) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	h.Bot.Send(msg)
}
