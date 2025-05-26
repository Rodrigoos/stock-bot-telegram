package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Ticker      string
	Quantity    int
	Price       float64
	PortfolioID uint
}
