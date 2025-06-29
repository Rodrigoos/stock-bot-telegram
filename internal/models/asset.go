package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Ticker        string
	Institution   string
	Quantity      int
	PurchasePrice float64
	Price         float64
	PortfolioID   uint
}
