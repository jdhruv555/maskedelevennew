.PHONY: help install dev build test clean docker-up docker-down docker-build docker-logs migrate seed

# Default target
help: ## Show this help message
	@echo "Masked 11 - Ecommerce Platform"
	@echo "=============================="
	@echo ""
	@echo "Available commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Development
install: ## Install all dependencies
	@echo "Installing dependencies..."
	cd backend && go mod tidy
	cd frontend && npm install

dev: ## Start development environment
	@echo "Starting development environment..."
	docker-compose -f infrastructure/docker/docker-compose.yml up -d
	@echo "Development environment started!"
	@echo "Frontend: http://localhost:3000"
	@echo "Backend API: http://localhost:8080"
	@echo "Grafana: http://localhost:3001 (admin/admin)"

dev-backend: ## Start only backend services
	@echo "Starting backend services..."
	docker-compose -f infrastructure/docker/docker-compose.yml up -d mongodb postgres redis backend

dev-frontend: ## Start only frontend
	@echo "Starting frontend..."
	cd frontend && npm run dev

# Database
migrate: ## Run database migrations
	@echo "Running database migrations..."
	cd backend && go run scripts/migrate.go

seed: ## Seed database with sample data
	@echo "Seeding database..."
	cd backend && go run scripts/seed.go

# Docker commands
docker-up: ## Start all Docker containers
	docker-compose -f infrastructure/docker/docker-compose.yml up -d

docker-down: ## Stop all Docker containers
	docker-compose -f infrastructure/docker/docker-compose.yml down

docker-build: ## Build all Docker images
	docker-compose -f infrastructure/docker/docker-compose.yml build

docker-logs: ## Show Docker logs
	docker-compose -f infrastructure/docker/docker-compose.yml logs -f

docker-clean: ## Clean Docker containers and volumes
	docker-compose -f infrastructure/docker/docker-compose.yml down -v
	docker system prune -f

# Testing
test: ## Run all tests
	@echo "Running backend tests..."
	cd backend && go test ./...
	@echo "Running frontend tests..."
	cd frontend && npm test

test-backend: ## Run backend tests only
	cd backend && go test ./...

test-frontend: ## Run frontend tests only
	cd frontend && npm test

# Building
build: ## Build both backend and frontend
	@echo "Building backend..."
	cd backend && go build -o bin/server cmd/server/main.go
	@echo "Building frontend..."
	cd frontend && npm run build

build-backend: ## Build backend only
	cd backend && go build -o bin/server cmd/server/main.go

build-frontend: ## Build frontend only
	cd frontend && npm run build

# Production
prod: ## Start production environment
	docker-compose -f infrastructure/docker/docker-compose.prod.yml up -d

prod-down: ## Stop production environment
	docker-compose -f infrastructure/docker/docker-compose.prod.yml down

# Linting and formatting
lint: ## Run linting
	@echo "Linting backend..."
	cd backend && golangci-lint run
	@echo "Linting frontend..."
	cd frontend && npm run lint

format: ## Format code
	@echo "Formatting backend..."
	cd backend && go fmt ./...
	@echo "Formatting frontend..."
	cd frontend && npm run format

# Monitoring
monitor: ## Open monitoring dashboards
	@echo "Opening monitoring dashboards..."
	@echo "Grafana: http://localhost:3001 (admin/admin)"
	@echo "Prometheus: http://localhost:9090"

logs: ## Show application logs
	@echo "Backend logs:"
	docker logs masked11_backend -f &
	@echo "Frontend logs:"
	docker logs masked11_frontend -f

# Database management
db-backup: ## Backup databases
	@echo "Backing up databases..."
	mkdir -p backups
	docker exec masked11_mongodb mongodump --out /data/backup
	docker cp masked11_mongodb:/data/backup backups/mongodb_$(shell date +%Y%m%d_%H%M%S)
	docker exec masked11_postgres pg_dump -U masked11_user masked11 > backups/postgres_$(shell date +%Y%m%d_%H%M%S).sql

db-restore: ## Restore databases (requires backup files)
	@echo "Restoring databases..."
	@echo "Please place backup files in the backups/ directory"

# Security
security-check: ## Run security checks
	@echo "Running security checks..."
	cd backend && gosec ./...
	cd frontend && npm audit

# Performance
benchmark: ## Run performance benchmarks
	@echo "Running performance benchmarks..."
	cd backend && go test -bench=. ./...

# Cleanup
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	cd backend && rm -rf bin/
	cd frontend && rm -rf .next/ node_modules/
	docker system prune -f

# Development helpers
shell-backend: ## Open shell in backend container
	docker exec -it masked11_backend sh

shell-frontend: ## Open shell in frontend container
	docker exec -it masked11_frontend sh

shell-mongodb: ## Open shell in MongoDB container
	docker exec -it masked11_mongodb mongosh

shell-postgres: ## Open shell in PostgreSQL container
	docker exec -it masked11_postgres psql -U masked11_user -d masked11

# API testing
api-test: ## Test API endpoints
	@echo "Testing API endpoints..."
	curl -f http://localhost:8080/health || echo "Backend not responding"
	curl -f http://localhost:3000 || echo "Frontend not responding"

# Documentation
docs: ## Generate documentation
	@echo "Generating documentation..."
	cd docs && make

# Quick start
quick-start: install docker-up migrate ## Quick start for development
	@echo "Masked 11 is ready!"
	@echo "Frontend: http://localhost:3000"
	@echo "Backend API: http://localhost:8080"
	@echo "API Docs: http://localhost:8080/docs"

# Status check
status: ## Check service status
	@echo "Checking service status..."
	@docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | grep masked11 || echo "No Masked 11 containers running"
