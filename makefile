APP_CONTAINER=stock_bot_app
DB_CONTAINER=stock_bot_db
DB_VOLUME=stock-bot-telegram_pgdata
DB_USER=stockuser
DB_NAME=stockdb

# Sobe o ambiente
up:
	docker compose up -d

# Derruba o ambiente
down:
	docker compose down

# Mostra os logs do banco
logs:
	docker compose logs -f db

# Entra no container do banco
psql:
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME)

# Reseta completamente o banco (cuidado! apaga tudo)
reset-db: down rm-volume up migrate

# Remove volume de dados
rm-volume:
	docker volume rm $(DB_VOLUME) || true

# Executa as migrations do GORM automaticamente
migrate:
	docker exec -it $(APP_CONTAINER) bash -c "go run ./cmd/migrate"

# Mostra volumes docker (debug)
volumes:
	docker volume ls
  	
# Criar carteiras com base em um csv	
seed:
	docker exec -it $(APP_CONTAINER) bash -c "go run ./cmd/seed"

seed-file:
	docker exec -it $(APP_CONTAINER) bash -c "go run ./cmd/seed/main.go $(FILE)"

# Status dos containers
status:
	docker ps

update-fund-prices:
	docker exec -it $(APP_CONTAINER) bash -c "go run ./cmd/update_fund_prices/main.go $(ID)"

update-stock-prices:
	docker exec -it $(APP_CONTAINER) bash -c "go run ./cmd/update_stock_prices/main.go $(ID)"

send-portfolio-message:
	docker exec -it $(APP_CONTAINER) bash -c "go run ./cmd/notify_telegram/main.go $(ID)"