# Welding API

A Go-based REST API with JWT authentication, built using clean architecture principles.

## Features

- User authentication with JWT
- PostgreSQL & MongoDB support
- Docker containerization
- Clean architecture
- Database migrations & seeding

## Project Structure

```
├── internal/
│   ├── cmd/                # Application entry point
│   ├── infrastructure/     # Infrastructure layer
│   └── module/             # Business modules
│       └── user/           # User module
├── pkg/                    # Shared packages
├── docker-compose.yaml     # Docker services
├── postman_collection.json # API documentation
└── Makefile                # Available commands
```

## Branching Strategy

### Core Branches
- `master` - Production code (protected, requires PR)
- `canary` - Development & integration branch
  - Always contains the latest stable changes
  - Recommended for API consumers and frontend development
  - Feature branches should be based on this branch

### Development Workflow
1. Start a new feature:
   ```bash
   git checkout canary
   git pull origin canary
   git checkout -b feature/name
   ```
2. Make and commit your changes
3. Push and create a pull request to `canary`
4. After review, merge into `canary`
5. For production releases, create a PR from `canary` to `master`

## Quick Start

### Prerequisites
- Go 1.25+
- Docker & Docker Compose

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/arfanxn/welding-golang.git
   cd welding-golang
   ```

2. **Setup environment and dependencies**
   ```bash
   make envs    # Create environment files from examples
   make deps    # Install Go dependencies
   ```

3. **Run migrations and seed data**
   ```bash
   make migrate-up
   make seed
   ```

4. **Start the server**
   ```bash
   make serve
   ```
   - Server runs on `http://localhost:8080` (customize via `APP_PORT` in `.env`)

### Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/arfanxn/welding-golang.git
   cd welding-golang
   ```

2. **Start services**
   ```bash
   # Start all services
   make docker-up
   
   # Or, to rebuild and start:
   # make docker-up-build
   ```

3. **Run migrations and seed data**
   ```bash
   make docker-migrate-up
   make docker-seed
   ```

**Access points:**
- **API**: `http://localhost:8080`
- **Nginx**: `http://localhost` (port 80)

## Available Commands

### Local Development
- `make build` - Build the application
- `make serve` - Start the application server
- `make migrate-up` - Run database migrations up (local)
- `make migrate-down` - Rollback database migrations (local)
- `make seed` - Seed database with sample data (local)

### Docker
- `make docker-up` - Start all services with Docker
- `make docker-up-build` - Rebuild and start all services
- `make docker-down` - Stop and remove all containers
- `make docker-restart` - Restart all services with rebuild
- `make docker-migrate-up` - Run database migrations (Docker)
- `make docker-migrate-down` - Rollback database migrations (Docker)
- `make docker-seed` - Seed database with sample data (Docker)
- `make docker-logs` - View container logs
- `make docker-ps` - List running containers

### Utilities
- `make envs` - Create environment files from examples
- `make deps` - Download Go module dependencies
- `make clean` - Clean build artifacts and Docker resources
- `make check-env` - Check if .env file exists
- `make check-env-docker` - Check if .env.docker file exists

## API Documentation

Import the Postman collection for complete API documentation:
- **File**: `postman_collection.json`
- **Base URL**: 
  - API Direct: `http://localhost:8080`
  - Via Nginx: `http://localhost` (port 80)

### Key Endpoints
- `GET /api/v1/health` - Health check
- `POST /api/v1/users/login` - User login
- `GET /api/v1/users/me` - Get user profile (protected)
- `DELETE /api/v1/users/logout` - User logout (protected)

### Default Credentials
- **Email**: `admin@gmail.com`
- **Password**: `11112222`

## License

MIT License
