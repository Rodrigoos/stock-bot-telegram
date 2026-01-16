package telegram

import (
	"fmt"
	"strings"

	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase"
	"github.com/Rodrigoos/stock-bot-telegram/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	Bot               *tgbotapi.BotAPI
	StartUseCase      *usecase.StartUseCase
	StockInfoUseCase  *usecase.StockInfoUseCase
	FundInfoUseCase   *usecase.FundInfoUseCase
	CriptoInfoUseCase *usecase.CriptoInfoUseCase
	PortfolioService  *usecase.PortfolioService
}

func NewHandler(
	bot *tgbotapi.BotAPI,
	startUC *usecase.StartUseCase,
	stockUC *usecase.StockInfoUseCase,
	fundUC *usecase.FundInfoUseCase,
	criptoUC *usecase.CriptoInfoUseCase,
	ps *usecase.PortfolioService,
) *Handler {
	return &Handler{
		Bot:               bot,
		StartUseCase:      startUC,
		StockInfoUseCase:  stockUC,
		FundInfoUseCase:   fundUC,
		CriptoInfoUseCase: criptoUC,
		PortfolioService:  ps,
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
			msg := "*📈 Stock Bot*\n" +
				"_Seu assistente de investimentos_\n\n" +
				"*Comandos disponíveis:*\n" +
				"• `/stock` ou `/acao` — Busca ações na B3\n" +
				"• `/fundo` — Busca fundos imobiliários\n" +
				"• `/cripto` — Busca criptomoedas\n" +
				"• `/portfolio` ou `/carteira` — Mostra sua carteira\n" +
				"• `/all-portfolios` — Lista todas as carteiras salvas\n" +
				"•by Rodrigo • github.com/Rodrigoos"

			message := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
			message.ParseMode = "Markdown"
			h.Bot.Send(message)

		case strings.HasPrefix(text, "/stock") || strings.HasPrefix(text, "/acao"):
			h.handleFindStock(update.Message)
		case strings.HasPrefix(text, "/fundo") || strings.HasPrefix(text, "/fund"):
			h.handleFindFund(update.Message)
		case strings.HasPrefix(text, "/cripto"):
			h.handleFindCripto(update.Message)
		case strings.HasPrefix(text, "/portfolio") || strings.HasPrefix(text, "/carteira"):
			h.handleGetPortfolio(update.Message)
		case strings.HasPrefix(text, "/all-portfolios") || strings.HasPrefix(text, "/totas-carteiras"):
			h.handleListPortfolios(update.Message)
		default:
			h.sendMessage(update.Message.Chat.ID, "Comando não reconhecido.", "Markdown")
		}
	}
}

func (h *Handler) handleFindStock(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Use o comando assim: /stock PETR4 ou /acao PETR4", "Text")
		return
	}

	ticker := strings.ToUpper(args[1])
	info, err := h.StockInfoUseCase.Execute(ticker)
	if err != nil {
		h.sendMessage(message.Chat.ID, fmt.Sprintf("Erro: %s", err), "Text")
		return
	}

	h.sendMessage(message.Chat.ID, info, "Text")
}

func (h *Handler) handleFindFund(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Use o comando assim: /fundo HGLG11 ou /fund HGLG11", "Text")
		return
	}

	ticker := strings.ToUpper(args[1])

	info, err := h.FundInfoUseCase.Execute(ticker)
	if err != nil {
		h.sendMessage(message.Chat.ID, "Erro ao buscar informação: "+err.Error(), "Text")
		return
	}

	h.sendMessage(message.Chat.ID, info, "Text")
}

func (h *Handler) handleFindCripto(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Use o comando assim: /cripto bitcoin", "Markdown")
		return
	}

	cripto := strings.ToUpper(args[1])

	info, err := h.CriptoInfoUseCase.Execute(cripto)
	if err != nil {
		h.sendMessage(message.Chat.ID, "Erro ao buscar informação: "+err.Error(), "Markdown")
		return
	}

	h.sendMessage(message.Chat.ID, info, "Text")
}

func (h *Handler) handleGetPortfolio(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Por favor, informe o nome da carteira. Exemplo: /carteira MinhaCarteira", "Text")
		return
	}

	portfolioName := args[1]
	portfolio, err := h.PortfolioService.GetPortfolioByName(portfolioName)
	if err != nil {
		h.sendMessage(message.Chat.ID, fmt.Sprintf("Erro: %s", err), "Markdown")
		return
	}

	msg := utils.FormatPortfolioMessage(*portfolio)

	h.sendMessage(message.Chat.ID, msg, "Markdown")
}

func (h *Handler) handleListPortfolios(message *tgbotapi.Message) {
	portfolios, err := h.PortfolioService.ListPortfolios()

	if err != nil {
		h.sendMessage(message.Chat.ID, fmt.Sprintf("Erro: %s", err), "Markdown")
		return
	}

	response := "Carteiras: \n"
	for _, portfolio := range portfolios {
		response += fmt.Sprintf("%s:  Total %s \n", portfolio.Name, utils.FormatBRL(portfolio.TotalValue()))
	}

	h.sendMessage(message.Chat.ID, response, "Markdown")
}

func (h *Handler) sendMessage(chatID int64, text string, parseMode string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = parseMode
	if parseMode == "Markdown" {
		msg.DisableWebPagePreview = true
	}
	if parseMode == "Text" {
		msg.ParseMode = ""
	}
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // Remove teclado virtual
	h.Bot.Send(msg)
}
