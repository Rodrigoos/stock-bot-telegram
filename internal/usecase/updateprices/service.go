// internal/usecase/updateprices/service.go

package updateprices

import (
	"fmt"
	"log"

	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
	"github.com/Rodrigoos/stock-bot-telegram/pkg/scraper"
	"gorm.io/gorm"
)

type Service struct {
	DB      *gorm.DB
	Scraper *scraper.StatusInvestScraper
}

func NewService(db *gorm.DB, scraper *scraper.StatusInvestScraper) *Service {
	return &Service{DB: db, Scraper: scraper}
}

func (s *Service) ExecuteStock(portfolioID uint) error {
	var portfolio models.Portfolio
	err := s.DB.Preload("Assets").First(&portfolio, portfolioID).Error
	if err != nil {
		return fmt.Errorf("portfólio não encontrado: %w", err)
	}

	for _, asset := range portfolio.Assets {
		price, err := s.Scraper.GetStockPrice(asset.Ticker)
		if err != nil {
			log.Printf("Erro ao buscar preço de %s: %v", asset.Ticker, err)
			continue
		}

		asset.Price = price
		if err := s.DB.Save(&asset).Error; err != nil {
			log.Printf("Erro ao atualizar ativo %s: %v", asset.Ticker, err)
		} else {
			log.Printf("Ativo %s atualizado para R$ %.2f", asset.Ticker, price)
		}
	}

	return nil
}

func (s *Service) ExecuteFund(portfolioID uint) error {
	var portfolio models.Portfolio
	err := s.DB.Preload("Assets").First(&portfolio, portfolioID).Error
	if err != nil {
		return fmt.Errorf("portfólio não encontrado: %w", err)
	}

	for _, asset := range portfolio.Assets {
		price, err := s.Scraper.GetFundPrice(asset.Ticker)
		if err != nil {
			log.Printf("Erro ao buscar preço de %s: %v", asset.Ticker, err)
			continue
		}

		asset.Price = price
		if err := s.DB.Save(&asset).Error; err != nil {
			log.Printf("Erro ao atualizar ativo %s: %v", asset.Ticker, err)
		} else {
			log.Printf("Ativo %s atualizado para R$ %.2f", asset.Ticker, price)
		}
	}

	return nil
}
