#!/bin/bash

# Masked 11 Ecommerce - Localhost Setup Script
# This script sets up the complete development environment

set -e

echo "🚀 Setting up Masked 11 Ecommerce for localhost development..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Docker is installed
check_docker() {
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed. Please install Docker first."
        exit 1
    fi
    print_success "Docker is installed"
}

# Check if Go is installed
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install Go 1.24.4+ first."
        exit 1
    fi
    print_success "Go is installed: $(go version)"
}

# Setup backend environment
setup_backend_env() {
    print_status "Setting up backend environment..."
    
    cd backend
    
    # Create .env file if it doesn't exist
    if [ ! -f .env ]; then
        cat > .env << 'EOF'
# =============================================================================
# Masked 11 Ecommerce Backend - Localhost Development Environment
# =============================================================================

# =============================================================================
# SERVER CONFIGURATION
# =============================================================================
APP_PORT=8080
APP_ENV=development
APP_NAME=Masked11-API
APP_VERSION=1.0.0

# =============================================================================
# DATABASE CONFIGURATION
# =============================================================================

# MongoDB Configuration
MONGO_URI=mongodb://localhost:27017
MONGO_DB=masked11
MONGO_USER=
MONGO_PASSWORD=
MONGO_AUTH_SOURCE=admin

# PostgreSQL Configuration
POSTGRES_URL=postgres://masked11_user:masked11_password@localhost:5432/masked11?sslmode=disable
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=masked11
POSTGRES_USER=masked11_user
POSTGRES_PASSWORD=masked11_password
POSTGRES_SSL_MODE=disable
POSTGRES_MAX_CONNECTIONS=100
POSTGRES_MIN_CONNECTIONS=5

# Redis Configuration
REDIS_URI=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
REDIS_POOL_SIZE=50
REDIS_MIN_IDLE_CONNS=10

# =============================================================================
# SECURITY CONFIGURATION
# =============================================================================
SESSION_SECRET=your-super-secret-key-change-in-production
JWT_SECRET=your-jwt-secret-key-change-in-production
JWT_EXPIRY=168h
BCRYPT_COST=12

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001,http://127.0.0.1:3000
ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
ALLOWED_HEADERS=Origin,Content-Type,Accept,Authorization,X-Requested-With

# Rate Limiting
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=1m

# =============================================================================
# ADMIN CONFIGURATION
# =============================================================================
ADMIN_EMAIL=admin@masked11.com
ADMIN_PASSWORD=admin123
ADMIN_NAME=Admin

# =============================================================================
# CACHE CONFIGURATION
# =============================================================================
CACHE_ENABLED=true
CACHE_TTL=3600
CACHE_MAX_SIZE=1000

# =============================================================================
# LOGGING CONFIGURATION
# =============================================================================
LOG_LEVEL=info
LOG_FORMAT=json
LOG_OUTPUT=stdout
LOG_FILE=

# =============================================================================
# MONITORING CONFIGURATION
# =============================================================================
METRICS_ENABLED=true
METRICS_PORT=9090
HEALTH_CHECK_ENABLED=true
PROFILING_ENABLED=false

# =============================================================================
# PERFORMANCE CONFIGURATION
# =============================================================================
GOMAXPROCS=1
GOGC=100
GODEBUG=netdns=go

# Database Connection Pool Settings
DB_MAX_OPEN_CONNS=100
DB_MAX_IDLE_CONNS=10
DB_CONN_MAX_LIFETIME=1h
DB_CONN_MAX_IDLE_TIME=30m

# =============================================================================
# DEVELOPMENT SETTINGS
# =============================================================================
DEBUG=true
HOT_RELOAD=true
CORS_DEBUG=false

# =============================================================================
# COMPRESSION SETTINGS
# =============================================================================
COMPRESSION_ENABLED=true
COMPRESSION_LEVEL=6
COMPRESSION_MIN_SIZE=1024

# =============================================================================
# TIMEOUT SETTINGS
# =============================================================================
READ_TIMEOUT=30s
WRITE_TIMEOUT=30s
IDLE_TIMEOUT=120s
SHUTDOWN_TIMEOUT=30s

# =============================================================================
# FEATURE FLAGS
# =============================================================================
FEATURE_CACHE=true
FEATURE_RATE_LIMITING=true
FEATURE_METRICS=true
FEATURE_HEALTH_CHECK=true
FEATURE_COMPRESSION=true
FEATURE_SECURITY_HEADERS=true
EOF
        print_success "Created backend .env file"
    else
        print_warning "Backend .env file already exists"
    fi
    
    cd ..
}

# Setup frontend environment
setup_frontend_env() {
    print_status "Setting up frontend environment..."
    
    cd frontend
    
    # Create .env.local file if it doesn't exist
    if [ ! -f .env.local ]; then
        cat > .env.local << 'EOF'
# =============================================================================
# Masked 11 Frontend - Localhost Development Environment
# =============================================================================

# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_APP_NAME=Masked 11
NEXT_PUBLIC_APP_VERSION=1.0.0

# Development Settings
NODE_ENV=development
NEXT_TELEMETRY_DISABLED=1

# Feature Flags
NEXT_PUBLIC_FEATURE_CART=true
NEXT_PUBLIC_FEATURE_SEARCH=true
NEXT_PUBLIC_FEATURE_FILTERS=true
NEXT_PUBLIC_FEATURE_REVIEWS=true
NEXT_PUBLIC_FEATURE_WISHLIST=true

# Analytics (Optional)
NEXT_PUBLIC_GA_ID=
NEXT_PUBLIC_GTM_ID=
EOF
        print_success "Created frontend .env.local file"
    else
        print_warning "Frontend .env.local file already exists"
    fi
    
    cd ..
}

# Start databases with Docker
start_databases() {
    print_status "Starting databases with Docker..."
    
    # Stop existing containers
    docker stop mongodb postgres redis 2>/dev/null || true
    docker rm mongodb postgres redis 2>/dev/null || true
    
    # Start MongoDB
    print_status "Starting MongoDB..."
    docker run -d \
        --name mongodb \
        -p 27017:27017 \
        -e MONGO_INITDB_DATABASE=masked11 \
        mongo:6.0
    
    # Start PostgreSQL
    print_status "Starting PostgreSQL..."
    docker run -d \
        --name postgres \
        -p 5432:5432 \
        -e POSTGRES_DB=masked11 \
        -e POSTGRES_USER=masked11_user \
        -e POSTGRES_PASSWORD=masked11_password \
        postgres:15
    
    # Start Redis
    print_status "Starting Redis..."
    docker run -d \
        --name redis \
        -p 6379:6379 \
        redis:7.0-alpine
    
    # Wait for databases to be ready
    print_status "Waiting for databases to be ready..."
    sleep 10
    
    print_success "All databases started successfully"
}

# Install backend dependencies
install_backend_deps() {
    print_status "Installing backend dependencies..."
    
    cd backend
    go mod download
    go mod tidy
    cd ..
    
    print_success "Backend dependencies installed"
}

# Install frontend dependencies
install_frontend_deps() {
    print_status "Installing frontend dependencies..."
    
    cd frontend
    npm install
    cd ..
    
    print_success "Frontend dependencies installed"
}

# Run database migrations
run_migrations() {
    print_status "Running database migrations..."
    
    cd backend
    go run scripts/migrate.go
    cd ..
    
    print_success "Database migrations completed"
}

# Start backend server
start_backend() {
    print_status "Starting backend server..."
    
    cd backend
    go run cmd/server/main.go &
    BACKEND_PID=$!
    cd ..
    
    print_success "Backend server started (PID: $BACKEND_PID)"
}

# Start frontend server
start_frontend() {
    print_status "Starting frontend server..."
    
    cd frontend
    npm run dev &
    FRONTEND_PID=$!
    cd ..
    
    print_success "Frontend server started (PID: $FRONTEND_PID)"
}

# Wait for servers to be ready
wait_for_servers() {
    print_status "Waiting for servers to be ready..."
    
    # Wait for backend
    while ! curl -s http://localhost:8080/health > /dev/null; do
        sleep 2
    done
    print_success "Backend is ready at http://localhost:8080"
    
    # Wait for frontend
    while ! curl -s http://localhost:3000 > /dev/null; do
        sleep 2
    done
    print_success "Frontend is ready at http://localhost:3000"
}

# Show status
show_status() {
    echo ""
    echo "🎉 Masked 11 Ecommerce is now running!"
    echo "======================================"
    echo ""
    echo "📱 Frontend: http://localhost:3000"
    echo "🔧 Backend API: http://localhost:8080"
    echo "🏥 Health Check: http://localhost:8080/health"
    echo "📊 Metrics: http://localhost:8080/metrics"
    echo ""
    echo "🗄️  Databases:"
    echo "   MongoDB: localhost:27017"
    echo "   PostgreSQL: localhost:5432"
    echo "   Redis: localhost:6379"
    echo ""
    echo "👤 Admin Account:"
    echo "   Email: admin@masked11.com"
    echo "   Password: admin123"
    echo ""
    echo "🛠️  Useful Commands:"
    echo "   Backend logs: cd backend && go run cmd/server/main.go"
    echo "   Frontend logs: cd frontend && npm run dev"
    echo "   Database logs: docker logs mongodb postgres redis"
    echo "   Stop all: pkill -f 'go run' && pkill -f 'npm run dev'"
    echo ""
}

# Main setup function
main() {
    echo "🚀 Masked 11 Ecommerce - Localhost Setup"
    echo "========================================"
    echo ""
    
    # Check prerequisites
    check_docker
    check_go
    
    # Setup environment files
    setup_backend_env
    setup_frontend_env
    
    # Start databases
    start_databases
    
    # Install dependencies
    install_backend_deps
    install_frontend_deps
    
    # Run migrations
    run_migrations
    
    # Start servers
    start_backend
    start_frontend
    
    # Wait for servers
    wait_for_servers
    
    # Show status
    show_status
    
    # Keep script running
    echo "Press Ctrl+C to stop all services..."
    trap 'echo ""; echo "🛑 Stopping services..."; pkill -f "go run"; pkill -f "npm run dev"; docker stop mongodb postgres redis; echo "✅ All services stopped"; exit 0' INT
    wait
}

# Run main function
main 