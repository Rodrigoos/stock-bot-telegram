# 🤖 stock-bot-telegram

Um bot do Telegram escrito em Go, com Clean Architecture, que permite consultar informações de ações e FIIs via comandos no chat. Atualmente, os dados são obtidos através de scraping do [StatusInvest](https://statusinvest.com.br) e outras fontes.

---

## 🚀 Funcionalidades

- Recebe comandos via Telegram
- Busca dados de ações e FIIs no StatusInvest
- Modular e escalável com Clean Architecture
- Persiste dados com PostgreSQL
- Configuração via .env
- Modular e escalável com Clean Architecture

---

## 📦 Instalação

Requer [Go 1.21+](https://go.dev/dl/)

```bash
git clone https://github.com/Rodrigoos/stock-bot-telegram.git
cd stock-bot-telegram
go mod tidy


stock-bot-telegram/
├── cmd/
│   └── bot/
│       └── main.go
├── internal/
│   ├── models/
│   │   ├── asset.go
│   │   └── portfolio.go
│   ├── infrastructure/
│   │   ├── database/
│   │   │   └── postgres.go
│   │   └── telegram/
│   │       └── bot.go
│   ├── interface/
│   │   └── telegram/
│   │       └── handler.go
│   ├── usecase/
│   │   ├── portfolio.go
│   │   └── scraper/
│   │       └── statusinvest.go
├── .env
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
└── README.md
