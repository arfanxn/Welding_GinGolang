# Welding - Makefile Commands
# ==========================

# ========== Auto execute commands ==========
ENV_FILE ?= .env
ifneq (,$(wildcard $(ENV_FILE)))
    include $(ENV_FILE)
    export
endif

.PHONY: help check-env check-env-docker clean setup build serve migrate-up migrate-down seed \
	docker-migrate-up docker-migrate-down docker-seed docker-up-build docker-up \
	docker-down docker-restart docker-logs docker-ps

# ========== Default target ==========
help:
	@echo "Welding - Available Commands:"
	@echo ""
	@echo "Local Development:"
	@echo "  make build           - Build the application"
	@echo "  make serve           - Start the application server"
	@echo "  make migrate-up      - Run database migrations up (local)"
	@echo "  make migrate-down    - Rollback database migrations (local)"
	@echo "  make seed            - Seed database with sample data (local)"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-up        - Start all services with Docker"
	@echo "  make docker-up-build  - Rebuild and start all services"
	@echo "  make docker-down      - Stop and remove all containers"
	@echo "  make docker-restart   - Restart all services with rebuild"
	@echo "  make docker-migrate-up   - Run database migrations (Docker)"
	@echo "  make docker-migrate-down - Rollback database migrations (Docker)"
	@echo "  make docker-seed      - Seed database with sample data (Docker)"
	@echo "  make docker-logs      - View container logs"
	@echo "  make docker-ps        - List running containers"
	@echo ""
	@echo "Utilities:"
	@echo "  make envs            - Create environment files from examples"
	@echo "  make deps            - Download Go module dependencies"
	@echo "  make clean           - Clean build artifacts and Docker resources"
	@echo "  make check-env       - Check if .env file exists"
	@echo "  make check-env-docker - Check if .env.docker file exists"


# ========== Utilities commands ==========
check-env:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "error: $(ENV_FILE) file not found. Run 'make setup' first."; \
		exit 1; \ 
	fi

check-env-docker:
	@if [ ! -f .env.docker ]; then \
		echo "error: .env.docker file not found. Run 'make setup' first."; \
		exit 1; \
	fi

clean:
	@if [ -f bin/welding ]; then rm bin/welding && echo "removed `bin/welding` binary"; fi
	docker system prune -f

# ========== Development commands ==========
envs:
	@if [ ! -f $(ENV_FILE) ]; then cp .env.example $(ENV_FILE) && echo "created $(ENV_FILE) file from .env.example"; fi
	@if [ ! -f .env.docker ]; then cp .env.example .env.docker && echo "created .env.docker file from .env.example"; fi

deps:
	go mod download

build:
	go build -o bin/welding main.go

serve:
	go run main.go serve

migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down

seed:
	go run main.go seed

# ========== Docker commands ==========
docker-up-build: check-env-docker
	docker compose up -d --build

docker-up: check-env-docker
	docker compose up -d

docker-down:
	docker compose down

docker-restart: docker-down docker-up-build

docker-migrate-up:
	docker compose exec welding ./server migrate up

docker-migrate-down:
	docker compose exec welding ./server migrate down

docker-seed:
	docker compose exec welding ./server seed

docker-logs:
	docker compose logs -f

docker-ps:
	docker compose ps