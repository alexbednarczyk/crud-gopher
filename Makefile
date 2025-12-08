.PHONY: run db-up db-down db-restart db-shell db-seed test lint

include .env
export

# Run the application locally
run:
	go run cmd/api/main.go

# Start the database container
db-up:
	docker compose up -d db

# Stop and remove the database container
db-down:
	docker compose down

# Restart the database container
db-restart: db-down db-up

# Access the database shell (psql)
db-shell:
	docker compose exec db psql -U $(POSTGRES_USER) -d $(POSTGRES_DATABASE)

# Seed the database with initial data
db-seed:
	docker compose exec -T db psql -U $(POSTGRES_USER) -d $(POSTGRES_DATABASE) < scripts/seed.sql

# Run tests
test:
	go test -v ./...

# Run linter
lint:
	golangci-lint run
