# Welding API

A Go-based REST API with JWT authentication, built using clean architecture principles.

## Features

- User authentication with JWT
- PostgreSQL & MongoDB support
- Docker containerization
- Clean architecture
- Database migrations & seeding

## Quick Start

### Prerequisites
- Go 1.25+
- Docker & Docker Compose

### Local Development

1. **Setup project**
   ```bash
   git clone https://github.com/arfanxn/welding-golang.git
   cd welding-golang
   make setup
   ```

2. **Start application**
   ```bash
   # Run migrations and seed data
   make migrate
   make seed
   
   # Start server
   make serve
   ```

   Server runs at `http://localhost:8080`

### Docker (Recommended)

1. **Start with Docker**
   ```bash
   git clone https://github.com/arfanxn/welding-golang.git
   cd welding-golang
   make setup
   make up
   ```

2. **Run migrations and seed data (Docker)**
   ```bash
   make docker-migrate
   make docker-seed
   ```

   **Access points:**
   - **API**: `http://localhost:8080`
   - **Nginx**: `http://localhost` (port 80)

## Available Commands

Run `make help` to see all available commands:

### Local Development
- `make setup` - Setup project (copy .env, install deps)
- `make build` - Build the application
- `make serve` - Start the application server
- `make migrate` - Run database migrations (local)
- `make seed` - Seed database with sample data (local)

### Docker
- `make up` - Start all services with Docker
- `make down` - Stop and remove all containers
- `make restart` - Restart all services
- `make docker-migrate` - Run database migrations (Docker)
- `make docker-seed` - Seed database with sample data (Docker)
- `make logs` - View container logs
- `make ps` - List running containers

### Utilities
- `make clean` - Clean build artifacts and Docker resources

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

## Project Structure

```
├── internal/
│   ├── infrastructure/     # Infrastructure layer
│   └── module/            # Business modules
│       └── user/          # User module
├── pkg/                   # Shared packages
├── docker-compose.yaml    # Docker services
├── postman_collection.json # API documentation
└── Makefile              # Available commands
```

## Environment Configuration

Copy `.env.example` to `.env` and modify as needed:
- Database connections
- JWT secret
- Application settings

## License

MIT License