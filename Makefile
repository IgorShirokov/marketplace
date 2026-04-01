DB_URL=postgres://postgres:12345678@localhost:9101/catalog-db-dev?sslmode=disable
MIGRATIONS_PATH=migrations/catalog

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-reset:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up
	
# 👇 принудительный сброс dirty состояния
migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(version)