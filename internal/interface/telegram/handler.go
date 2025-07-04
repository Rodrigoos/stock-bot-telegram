package telegram

import (
	"fmt"
	"strings"

	"github.com/Rodrigoos/stock-bot-telegram/internal/usecase"
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
			msg := h.StartUseCase.Execute()
			// log.Println(update.Message.Chat.ID)
			h.sendMessage(update.Message.Chat.ID, msg)
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
			h.sendMessage(update.Message.Chat.ID, "Comando não reconhecido.")
		}
	}
}

func (h *Handler) handleFindStock(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Use o comando assim: /stock PETR4 ou /acao PETR4")
		return
	}

	ticker := strings.ToUpper(args[1])
	info, err := h.StockInfoUseCase.Execute(ticker)
	if err != nil {
		h.sendMessage(message.Chat.ID, fmt.Sprintf("Erro: %s", err))
		return
	}

	h.sendMessage(message.Chat.ID, info)
}

func (h *Handler) handleFindFund(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Use o comando assim: /fundo HGLG11 ou /fund HGLG11")
		return
	}

	ticker := strings.ToUpper(args[1])

	info, err := h.FundInfoUseCase.Execute(ticker)
	if err != nil {
		h.sendMessage(message.Chat.ID, "Erro ao buscar informação: "+err.Error())
		return
	}

	h.sendMessage(message.Chat.ID, info)
}

func (h *Handler) handleFindCripto(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Use o comando assim: /cripto bitcoin")
		return
	}

	cripto := strings.ToUpper(args[1])

	info, err := h.CriptoInfoUseCase.Execute(cripto)
	if err != nil {
		h.sendMessage(message.Chat.ID, "Erro ao buscar informação: "+err.Error())
		return
	}

	h.sendMessage(message.Chat.ID, info)
}

func (h *Handler) handleGetPortfolio(message *tgbotapi.Message) {
	args := strings.SplitN(message.Text, " ", 2)
	if len(args) < 2 {
		h.sendMessage(message.Chat.ID, "Por favor, informe o nome da carteira. Exemplo: /carteira MinhaCarteira")
		return
	}

	portfolioName := args[1]
	portfolio, err := h.PortfolioService.GetPortfolioByName(portfolioName)
	if err != nil {
		h.sendMessage(message.Chat.ID, fmt.Sprintf("Erro: %s", err))
		return
	}

	response := fmt.Sprintf("Carteira: %s\nAtivos:\n", portfolio.Name)
	for _, asset := range portfolio.Assets {
		response += fmt.Sprintf("%s:\n%d x R$%.2f\nTotal R$%.2f\n",
			asset.Ticker, asset.Quantity, asset.Price, (float64(asset.Quantity) * asset.Price))
	}

	response += fmt.Sprintf("Total de Ativos: %d\n", len(portfolio.Assets))
	response += fmt.Sprintf("Quantidade Total de Ativos: %d\n", portfolio.TotalQuantity())
	response += fmt.Sprintf("Total em Carteira: R$%.2f\n", portfolio.TotalValue())

	h.sendMessage(message.Chat.ID, response)
}

func (h *Handler) handleListPortfolios(message *tgbotapi.Message) {
	portfolios, err := h.PortfolioService.ListPortfolios()

	if err != nil {
		h.sendMessage(message.Chat.ID, fmt.Sprintf("Erro: %s", err))
		return
	}

	response := fmt.Sprintf("Carteiras: \n")
	for _, portfolio := range portfolios {
		response += fmt.Sprintf("%s:  Total R$%.2f \n", portfolio.Name, portfolio.TotalValue())
	}

	h.sendMessage(message.Chat.ID, response)
}

func (h *Handler) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	h.Bot.Send(msg)
}
