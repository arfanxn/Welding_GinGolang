# Welding - Makefile Commands
# ==========================

# Load environment variables
ENV_FILE ?= .env
ifneq (,$(wildcard $(ENV_FILE)))
    include $(ENV_FILE)
    export
endif

.PHONY: help setup build serve migrate seed docker-migrate docker-seed up down restart logs ps clean

# Default target
help:
	@echo "Welding - Available Commands:"
	@echo ""
	@echo "Local Development:"
	@echo "  make setup     - Setup project (copy .env, install deps)"
	@echo "  make build     - Build the application"
	@echo "  make serve     - Start the application server"
	@echo "  make migrate   - Run database migrations (local)"
	@echo "  make seed      - Seed database with sample data (local)"
	@echo ""
	@echo "Docker:"
	@echo "  make up            - Start all services with Docker"
	@echo "  make down          - Stop and remove all containers"
	@echo "  make restart       - Restart all services"
	@echo "  make docker-migrate - Run database migrations (Docker)"
	@echo "  make docker-seed   - Seed database with sample data (Docker)"
	@echo "  make logs          - View container logs"
	@echo "  make ps            - List running containers"
	@echo ""
	@echo "Utilities:"
	@echo "  make clean     - Clean build artifacts and Docker resources"

# Development Commands
setup:
	@echo "Setting up project..."
	@if [ ! -f .env ]; then cp .env.example .env && echo "Created .env file from .env.example"; fi
	@if [ ! -f .env.docker ]; then cp .env.example .env.docker && echo "Created .env.docker file"; fi
	go mod download
	@echo "Setup complete!"

build:
	@echo "Building application..."
	go build -o bin/welding main.go
	@echo "Build complete!"

serve:
	@echo "Starting application server..."
	go run main.go serve

migrate:
	@echo "Running database migrations..."
	go run main.go migrate

seed:
	@echo "Seeding database..."
	go run main.go seed

# Docker Commands
docker-migrate:
	@echo "Running database migrations in Docker..."
	docker compose exec welding ./server migrate

docker-seed:
	@echo "Seeding database in Docker..."
	docker compose exec welding ./server seed

# Docker Service Commands
check-env:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "Error: $(ENV_FILE) file not found. Run 'make setup' first."; \
		exit 1; \
	fi

up: check-env
	@echo "Starting services with Docker..."
	docker compose up -d --build
	@echo "Services started!"
	@echo "  - API: http://localhost:8080"
	@echo "  - Nginx: http://localhost (port 80)"
down:
	@echo "Stopping services..."
	docker compose down

restart: down up

logs:
	@echo "Viewing container logs (Ctrl+C to exit)..."
	docker compose logs -f

ps:
	@echo "Container status:"
	docker compose ps

# Utility Commands
clean:
	@echo "Cleaning up..."
	@if [ -f bin/welding ]; then rm bin/welding && echo "Removed binary"; fi
	@docker system prune -f
	@echo "Cleanup complete!"
