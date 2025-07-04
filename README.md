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
```

## 🛠️ Comandos do Makefile

O projeto possui um `Makefile` para facilitar tarefas comuns com Docker. Veja os principais comandos:

```bash
# Sobe o ambiente (containers em background)
make up

# Derruba o ambiente (para e remove containers)
make down

# Mostra os logs do banco de dados
make logs

# Acessa o banco de dados via psql
make psql

# Reseta completamente o banco de dados (remove tudo e recria)
make reset-db

# Remove o volume de dados do banco (apaga todos os dados)
make rm-volume

# Executa as migrations do banco de dados
make migrate

# Mostra os volumes Docker existentes
make volumes

# Popula o banco usando um arquivo específico
make seed-file FILE=path/do/arquivo.csv

# Mostra o status dos containers Docker
make status

# Atualiza os preços dos fundos (pode passar ID opcional: make update-fund-prices ID=123)
make update-fund-prices

# Atualiza os preços das ações (pode passar ID opcional: make update-stock-prices ID=123)
make update-stock-prices
```

Para ver todos os comandos disponíveis, rode:

```bash
make help
```

