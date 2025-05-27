package usecase

import (
	"fmt"

	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
	"gorm.io/gorm"
)

type PortfolioService struct {
	DB *gorm.DB
}

func NewPortfolioService(db *gorm.DB) *PortfolioService {
	return &PortfolioService{DB: db}
}

// Cria uma nova carteira
func (s *PortfolioService) CreatePortfolio(name string) (*models.Portfolio, error) {
	portfolio := &models.Portfolio{Name: name}
	if err := s.DB.Create(portfolio).Error; err != nil {
		return nil, fmt.Errorf("erro ao criar carteira: %w", err)
	}
	return portfolio, nil
}

// Adiciona um ativo a uma carteira
func (s *PortfolioService) AddAssetToPortfolio(portfolioID uint, asset *models.Asset) error {
	asset.PortfolioID = portfolioID
	if err := s.DB.Create(asset).Error; err != nil {
		return fmt.Errorf("erro ao adicionar ativo: %w", err)
	}
	return nil
}

// Busca carteira por nome, junto com os ativos
func (s *PortfolioService) GetPortfolioByName(name string) (*models.Portfolio, error) {
	var portfolio models.Portfolio
	if err := s.DB.Preload("Assets").Where("name = ?", name).First(&portfolio).Error; err != nil {
		return nil, fmt.Errorf("carteira não encontrada: %w", err)
	}
	return &portfolio, nil
}

// Lista todas as carteiras
func (s *PortfolioService) ListPortfolios() ([]models.Portfolio, error) {
	var portfolios []models.Portfolio
	if err := s.DB.Preload("Assets").Find(&portfolios).Error; err != nil {
		return nil, fmt.Errorf("erro ao listar carteiras: %w", err)
	}
	return portfolios, nil
}
