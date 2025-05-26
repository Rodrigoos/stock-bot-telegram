# 🤖 stock-bot-telegram

Um bot do Telegram escrito em Go, com Clean Architecture, que permite consultar informações de ações e FIIs via comandos no chat. Atualmente, os dados são obtidos através de scraping do [StatusInvest](https://statusinvest.com.br) e outras fontes.

---

## 🚀 Funcionalidades

- Comando `/start`: mensagem de boas-vindas
- Comando `/quote [TICKER]`: retorna a cotação da ação ou FII
- Modular e escalável com Clean Architecture
- Scrapers desacoplados para diferentes fontes de dados

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
│       └── main.go                  // Entrada principal do bot
│
├── internal/
│   ├── interface/
│   │   └── telegram/
│   │       └── handler.go          // Lida com comandos do Telegram
│
│   ├── infrastructure/
│   │   └── telegram/
│   │       └── bot.go              // Conexão com o Telegram
│   │
│   └── usecase/
│       ├── start/
│       │   └── start.go          // Lógica do /start
│       │
│       └── stockinfo/
│           └── stockinfo.go          // Lógica para buscar cotação
│
├── pkg/
│   └── scraper/
│       └── statusinvest.go         // Scraper para StatusInvest
│
├── go.mod
└── go.sum
