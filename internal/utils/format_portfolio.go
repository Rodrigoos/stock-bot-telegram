package utils

import (
	"fmt"
	"strings"

	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
)

func FormatPortfolioMessage(p models.Portfolio) string {
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
			FormatBRL(asset.Price),
			FormatBRL(total),
		))
		totalGeral += total
	}

	b.WriteString(strings.Repeat("-", 42) + "\n")
	b.WriteString(fmt.Sprintf("Total geral: %s\n", FormatBRL(totalGeral)))
	b.WriteString("```")

	return b.String()
}
