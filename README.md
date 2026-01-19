# 🤖 stock-bot-telegram

Um bot do Telegram escrito em Go, com Clean Architecture, que permite consultar informações de ações e FIIs via comandos no chat. Atualmente, os dados são obtidos através de scraping do [StatusInvest](https://statusinvest.com.br) e outras fontes.

---

## 🚀 Funcionalidades

- Recebe comandos via Telegram
- Busca dados de ações e FIIs no StatusInvest
- Modular e escalável com Clean Architecture
- Persiste dados com PostgreSQL
- Configuração via .env
  
---

## 📦 Instalação

Requer [Go 1.21+](https://go.dev/dl/)

```bash
git clone https://github.com/Rodrigoos/stock-bot-telegram.git
cd stock-bot-telegram
go mod tidy


stock-bot-telegram/
├── cmd/
│   ├── bot/
│   │   └── main.go
│   ├── migrate/
│   │   └── main.go
│   ├── notify_portfolio/
│   │   └── main.go
│   ├── seed/
│   │   └── main.go
│   ├── update_fund_prices/
│   │   └── main.go
│   └── update_stock_price/
│       └── main.go
│
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

## 🖼️ Capturas de telas

### start
<img width="965" height="876" alt="image" src="https://github.com/user-attachments/assets/8239018f-ff8e-49ad-9c06-bff8c8062459" />


### seed file 
<img width="685" height="394" alt="image" src="https://github.com/user-attachments/assets/faedd3f7-501f-4ed5-b62b-a25238945e3d" />

### stock
<img width="742" height="138" alt="image" src="https://github.com/user-attachments/assets/c645787a-39af-4d33-9995-278026f5542d" />

### carteira

<img width="748" height="638" alt="image" src="https://github.com/user-attachments/assets/86041814-6216-4ea9-b7a6-ed1483a8b5fa" />


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


