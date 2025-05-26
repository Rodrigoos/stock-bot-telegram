// package usecase

// import (
// 	"log"

// 	"github.com/Rodrigoos/stock-bot-telegram/internal/infrastructure/db"
// 	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
// )

// // AddToPortfolio salva uma ação ou fundo imobiliário na carteira
// func AddToPortfolio(ticker string, quantity int, price float64) {
// 	portfolio := models.Portfolio{Ticker: ticker, Quantity: quantity, Price: price}
// 	result := db.DB.Create(&portfolio)
// 	if result.Error != nil {
// 		log.Println("Erro ao adicionar à carteira:", result.Error)
// 	} else {
// 		log.Println("Ação adicionada à carteira com sucesso!")
// 	}
// }
