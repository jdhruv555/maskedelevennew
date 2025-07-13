# Masked 11 Backend Makefile
# Professional-level build and deployment automation

# Variables
APP_NAME := masked11-backend
APP_VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GO_VERSION := $(shell go version | awk '{print $$3}')

# Build flags
LDFLAGS := -w -s \
	-X main.Version=$(APP_VERSION) \
	-X main.BuildTime=$(BUILD_TIME) \
	-X main.GoVersion=$(GO_VERSION)

# Directories
BIN_DIR := bin
DIST_DIR := dist
TEST_DIR := test
COVERAGE_DIR := coverage

# Files
MAIN_FILE := cmd/server/main.go
BINARY := $(BIN_DIR)/server
DOCKER_IMAGE := $(APP_NAME):$(APP_VERSION)
DOCKER_IMAGE_LATEST := $(APP_NAME):latest

# Environment
ENV_FILE := .env
ENV_EXAMPLE := .env.example

# Database URLs
MONGO_URL := mongodb://localhost:27017
POSTGRES_URL := postgres://masked11_user:masked11_password@localhost:5432/masked11?sslmode=disable
REDIS_URL := localhost:6379

# Test settings
TEST_TIMEOUT := 30s
COVERAGE_THRESHOLD := 80

# Docker settings
DOCKER_COMPOSE_FILE := ../infrastructure/docker/docker-compose.yml
DOCKER_COMPOSE_PROD_FILE := ../infrastructure/docker/docker-compose.prod.yml

# Default target
.PHONY: help
help: ## Show this help message
	@echo "Masked 11 Backend - Available Commands"
	@echo "======================================"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# =============================================================================
# DEVELOPMENT COMMANDS
# =============================================================================

.PHONY: dev
dev: ## Start development server with hot reload
	@echo "🚀 Starting development server..."
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "⚠️  .env file not found. Creating from example..."; \
		cp $(ENV_EXAMPLE) $(ENV_FILE); \
	fi
	@go run $(MAIN_FILE)

.PHONY: build
build: ## Build the application
	@echo "🔨 Building $(APP_NAME) v$(APP_VERSION)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-ldflags="$(LDFLAGS)" \
		-o $(BINARY) \
		$(MAIN_FILE)
	@echo "✅ Build complete: $(BINARY)"

.PHONY: build-all
build-all: ## Build for multiple platforms
	@echo "🔨 Building for multiple platforms..."
	@mkdir -p $(DIST_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-ldflags="$(LDFLAGS)" \
		-o $(DIST_DIR)/$(APP_NAME)-linux-amd64 \
		$(MAIN_FILE)
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags="$(LDFLAGS)" \
		-o $(DIST_DIR)/$(APP_NAME)-darwin-amd64 \
		$(MAIN_FILE)
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
		-ldflags="$(LDFLAGS)" \
		-o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe \
		$(MAIN_FILE)
	@echo "✅ Multi-platform build complete"

.PHONY: clean
clean: ## Clean build artifacts
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BIN_DIR) $(DIST_DIR) $(COVERAGE_DIR)
	@go clean -cache -testcache
	@echo "✅ Clean complete"

# =============================================================================
# TESTING COMMANDS
# =============================================================================

.PHONY: test
test: ## Run all tests
	@echo "🧪 Running tests..."
	@go test -v -timeout=$(TEST_TIMEOUT) ./...

.PHONY: test-short
test-short: ## Run tests with short timeout
	@echo "🧪 Running quick tests..."
	@go test -v -short -timeout=10s ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage report
	@echo "🧪 Running tests with coverage..."
	@mkdir -p $(COVERAGE_DIR)
	@go test -v -timeout=$(TEST_TIMEOUT) -coverprofile=$(COVERAGE_DIR)/coverage.out ./...
	@go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@go tool cover -func=$(COVERAGE_DIR)/coverage.out
	@echo "📊 Coverage report: $(COVERAGE_DIR)/coverage.html"

.PHONY: test-benchmark
test-benchmark: ## Run benchmark tests
	@echo "🏃 Running benchmarks..."
	@go test -v -bench=. -benchmem ./...

.PHONY: test-race
test-race: ## Run tests with race detection
	@echo "🏁 Running tests with race detection..."
	@go test -v -race -timeout=$(TEST_TIMEOUT) ./...

# =============================================================================
# DATABASE COMMANDS
# =============================================================================

.PHONY: db-setup
db-setup: ## Setup all databases
	@echo "🗄️  Setting up databases..."
	@echo "📦 MongoDB..."
	@docker run -d -p 27017:27017 --name mongodb mongo:6.0 || echo "MongoDB already running"
	@echo "🐘 PostgreSQL..."
	@docker run -d -p 5432:5432 --name postgres \
		-e POSTGRES_DB=masked11 \
		-e POSTGRES_USER=masked11_user \
		-e POSTGRES_PASSWORD=masked11_password \
		postgres:15 || echo "PostgreSQL already running"
	@echo "🔴 Redis..."
	@docker run -d -p 6379:6379 --name redis redis:7.0-alpine || echo "Redis already running"
	@echo "✅ Database setup complete"

.PHONY: db-migrate
db-migrate: ## Run database migrations
	@echo "🔄 Running database migrations..."
	@go run scripts/migrate.go

.PHONY: db-reset
db-reset: ## Reset all databases
	@echo "🔄 Resetting databases..."
	@docker stop mongodb postgres redis 2>/dev/null || true
	@docker rm mongodb postgres redis 2>/dev/null || true
	@make db-setup
	@make db-migrate

.PHONY: db-logs
db-logs: ## Show database logs
	@echo "📋 Database logs..."
	@docker logs mongodb --tail=20
	@docker logs postgres --tail=20
	@docker logs redis --tail=20

# =============================================================================
# DOCKER COMMANDS
# =============================================================================

.PHONY: docker-build
docker-build: ## Build Docker image
	@echo "🐳 Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .
	@docker tag $(DOCKER_IMAGE) $(DOCKER_IMAGE_LATEST)
	@echo "✅ Docker image built: $(DOCKER_IMAGE)"

.PHONY: docker-run
docker-run: ## Run Docker container
	@echo "🐳 Running Docker container..."
	@docker run -d --name $(APP_NAME) \
		-p 8080:8080 \
		--env-file $(ENV_FILE) \
		$(DOCKER_IMAGE)

.PHONY: docker-stop
docker-stop: ## Stop Docker container
	@echo "🛑 Stopping Docker container..."
	@docker stop $(APP_NAME) 2>/dev/null || true
	@docker rm $(APP_NAME) 2>/dev/null || true

.PHONY: docker-logs
docker-logs: ## Show Docker logs
	@docker logs -f $(APP_NAME)

.PHONY: docker-shell
docker-shell: ## Shell into Docker container
	@docker exec -it $(APP_NAME) /bin/sh

# =============================================================================
# DOCKER COMPOSE COMMANDS
# =============================================================================

.PHONY: compose-up
compose-up: ## Start all services with Docker Compose
	@echo "🐳 Starting all services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: compose-down
compose-down: ## Stop all services
	@echo "🛑 Stopping all services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: compose-logs
compose-logs: ## Show Docker Compose logs
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

.PHONY: compose-build
compose-build: ## Build and start all services
	@echo "🔨 Building and starting all services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up --build -d

# =============================================================================
# PRODUCTION COMMANDS
# =============================================================================

.PHONY: prod-build
prod-build: ## Build production Docker image
	@echo "🏭 Building production image..."
	@docker build --target production -t $(DOCKER_IMAGE) .

.PHONY: prod-deploy
prod-deploy: ## Deploy to production
	@echo "🚀 Deploying to production..."
	@docker-compose -f $(DOCKER_COMPOSE_PROD_FILE) up -d

.PHONY: prod-logs
prod-logs: ## Show production logs
	@docker-compose -f $(DOCKER_COMPOSE_PROD_FILE) logs -f

# =============================================================================
# MONITORING COMMANDS
# =============================================================================

.PHONY: health-check
health-check: ## Check application health
	@echo "🏥 Health check..."
	@curl -f http://localhost:8080/health || echo "❌ Health check failed"

.PHONY: metrics
metrics: ## Show application metrics
	@echo "📊 Application metrics..."
	@curl -s http://localhost:8080/metrics | head -20

.PHONY: performance-test
performance-test: ## Run performance tests
	@echo "⚡ Running performance tests..."
	@if command -v ab >/dev/null 2>&1; then \
		echo "Apache Bench test:"; \
		ab -n 1000 -c 10 http://localhost:8080/api/products; \
	else \
		echo "Apache Bench not found. Install with: brew install httpd"; \
	fi

# =============================================================================
# UTILITY COMMANDS
# =============================================================================

.PHONY: deps
deps: ## Download dependencies
	@echo "📦 Downloading dependencies..."
	@go mod download
	@go mod tidy

.PHONY: lint
lint: ## Run linter
	@echo "🔍 Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

.PHONY: fmt
fmt: ## Format code
	@echo "🎨 Formatting code..."
	@go fmt ./...

.PHONY: vet
vet: ## Run go vet
	@echo "🔍 Running go vet..."
	@go vet ./...

.PHONY: generate
generate: ## Generate code (if any)
	@echo "⚙️  Generating code..."
	@go generate ./...

.PHONY: swagger
swagger: ## Generate API documentation
	@echo "📚 Generating API documentation..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g $(MAIN_FILE); \
	else \
		echo "swag not found. Install with: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# =============================================================================
# DEVELOPMENT TOOLS
# =============================================================================

.PHONY: install-tools
install-tools: ## Install development tools
	@echo "🛠️  Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/go-delve/delve/cmd/dlv@latest
	@echo "✅ Tools installed"

.PHONY: debug
debug: ## Start debug server
	@echo "🐛 Starting debug server..."
	@dlv debug $(MAIN_FILE) --listen=:2345 --headless=true --api-version=2 --accept-multiclient

.PHONY: profile
profile: ## Run with profiling
	@echo "📊 Starting with profiling..."
	@go run $(MAIN_FILE) -cpuprofile=cpu.prof -memprofile=mem.prof

# =============================================================================
# SECURITY COMMANDS
# =============================================================================

.PHONY: security-scan
security-scan: ## Run security scan
	@echo "🔒 Running security scan..."
	@if command -v trivy >/dev/null 2>&1; then \
		trivy fs --exit-code 1 --severity HIGH,CRITICAL .; \
	else \
		echo "trivy not found. Install with: brew install trivy"; \
	fi

.PHONY: audit
audit: ## Audit dependencies
	@echo "🔍 Auditing dependencies..."
	@go list -json -deps . | nancy sleuth

# =============================================================================
# BACKUP AND RESTORE
# =============================================================================

.PHONY: backup
backup: ## Backup databases
	@echo "💾 Creating database backups..."
	@mkdir -p backups
	@docker exec mongodb mongodump --out /dump
	@docker cp mongodb:/dump backups/mongodb-$(shell date +%Y%m%d_%H%M%S)
	@docker exec postgres pg_dump -U masked11_user masked11 > backups/postgres-$(shell date +%Y%m%d_%H%M%S).sql
	@echo "✅ Backups created in backups/"

.PHONY: restore
restore: ## Restore databases (requires backup files)
	@echo "🔄 Restoring databases..."
	@echo "Please specify backup files to restore"

# =============================================================================
# CLEANUP COMMANDS
# =============================================================================

.PHONY: cleanup
cleanup: ## Clean up all containers and images
	@echo "🧹 Cleaning up..."
	@docker stop $$(docker ps -aq) 2>/dev/null || true
	@docker rm $$(docker ps -aq) 2>/dev/null || true
	@docker rmi $$(docker images -q) 2>/dev/null || true
	@docker system prune -f
	@echo "✅ Cleanup complete"

# =============================================================================
# INFORMATION COMMANDS
# =============================================================================

.PHONY: info
info: ## Show build information
	@echo "📋 Build Information"
	@echo "==================="
	@echo "App Name: $(APP_NAME)"
	@echo "Version: $(APP_VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Main File: $(MAIN_FILE)"
	@echo "Binary: $(BINARY)"

.PHONY: status
status: ## Show system status
	@echo "📊 System Status"
	@echo "================"
	@echo "Docker containers:"
	@docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
	@echo ""
	@echo "Database connections:"
	@curl -s http://localhost:8080/health 2>/dev/null || echo "❌ Backend not running"

# =============================================================================
# ALIASES
# =============================================================================

.PHONY: run
run: dev ## Alias for dev

.PHONY: test-all
test-all: test test-coverage test-race ## Run all tests

.PHONY: build-test
build-test: build test ## Build and test

.PHONY: setup
setup: deps db-setup db-migrate ## Complete setup

.PHONY: start
start: compose-up ## Start all services

.PHONY: stop
stop: compose-down ## Stop all services

# =============================================================================
# SPECIAL TARGETS
# =============================================================================

.PHONY: .PHONY
.PHONY: all
all: clean deps build test ## Complete build pipeline

.PHONY: ci
ci: deps lint test-coverage build ## CI pipeline

.PHONY: release
release: clean build-all test-all docker-build ## Release pipeline 