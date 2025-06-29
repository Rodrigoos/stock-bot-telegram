package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/db"
	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: .env não encontrado, seguindo com variáveis de ambiente do sistema")
	}

	err := seedFromXPCSV("cmd/seed/csv/acoes.csv", "Acoes")
	if err != nil {
		log.Fatalf("Erro ao semear dados: %v", err)
	}

	fmt.Println("Seed concluído com sucesso.")
}

func seedFromXPCSV(path string, portfolioName string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	db.Connect()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // permite quantidade variável de colunas
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("erro ao ler csv: %w", err)
	}

	if len(records) < 1 {
		return fmt.Errorf("csv vazio")
	}

	var portfolio models.Portfolio
	err = db.DB.FirstOrCreate(&portfolio, models.Portfolio{Name: portfolioName}).Error
	if err != nil {
		return fmt.Errorf("erro ao criar portfolio: %w", err)
	}

	for _, row := range records[1:] {
		if len(row) < 14 {
			continue // pula linhas inválidas ou vazias
		}

		institution := strings.TrimSpace(row[1])
		ticker := strings.TrimSpace(row[3])
		quantStr := strings.TrimSpace(row[8])
		priceStr := strings.TrimSpace(row[12])

		// Validação básica
		if ticker == "" || quantStr == "" || priceStr == "" {
			continue
		}

		quantidade, err := strconv.Atoi(quantStr)
		if err != nil || quantidade == 0 {
			continue
		}

		// Remove "R$" e converte vírgula para ponto
		priceStr = strings.ReplaceAll(priceStr, "R$", "")
		priceStr = strings.ReplaceAll(priceStr, ",", ".")
		priceStr = strings.Trim(priceStr, "\" ")

		preco, err := strconv.ParseFloat(priceStr, 64)
		if err != nil || preco == 0 {
			continue
		}
		log.Printf("Preco %f", preco)

		asset := models.Asset{
			Ticker:      ticker,
			Institution: institution,
			Quantity:    quantidade,
			Price:       preco,
			PortfolioID: portfolio.ID,
		}

		if err := db.DB.Create(&asset).Error; err != nil {
			log.Printf("Erro ao criar ativo %s: %v", ticker, err)
		}
	}

	return nil
}
