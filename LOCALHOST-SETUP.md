# 🚀 Masked 11 Ecommerce - Localhost Setup Guide

This guide will help you set up the complete Masked 11 ecommerce application on your localhost.

## 📋 Prerequisites

Before starting, make sure you have the following installed:

- **Docker** - For running databases
- **Go 1.24.4+** - For the backend
- **Node.js 18+** - For the frontend
- **Git** - For version control

## 🚀 Quick Start (Automated)

### Option 1: Use the Setup Script

1. **Make the script executable:**
   ```bash
   chmod +x setup-localhost.sh
   ```

2. **Run the setup script:**
   ```bash
   ./setup-localhost.sh
   ```

3. **Open your browser:**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

## 🛠️ Manual Setup

### Step 1: Environment Setup

#### Backend Environment
Create `backend/.env` file:

```bash
cd backend
```

Create the `.env` file with the following content:

```env
# Server Configuration
APP_PORT=8080
APP_ENV=development
APP_NAME=Masked11-API
APP_VERSION=1.0.0

# Database Configuration
MONGO_URI=mongodb://localhost:27017
POSTGRES_URL=postgres://masked11_user:masked11_password@localhost:5432/masked11?sslmode=disable
REDIS_URI=localhost:6379

# Security Configuration
JWT_SECRET=your-super-secret-jwt-key
SESSION_SECRET=your-super-secret-session-key

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001,http://127.0.0.1:3000
ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
ALLOWED_HEADERS=Origin,Content-Type,Accept,Authorization,X-Requested-With

# Rate Limiting
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=1m

# Admin Configuration
ADMIN_EMAIL=admin@masked11.com
ADMIN_PASSWORD=admin123
ADMIN_NAME=Admin

# Cache Configuration
CACHE_ENABLED=true
CACHE_TTL=3600
CACHE_MAX_SIZE=1000

# Performance Configuration
GOMAXPROCS=1
GOGC=100
GODEBUG=netdns=go

# Development Settings
DEBUG=true
HOT_RELOAD=true
CORS_DEBUG=false
```

#### Frontend Environment
Create `frontend/.env.local` file:

```bash
cd frontend
```

Create the `.env.local` file with the following content:

```env
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
```

### Step 2: Start Databases

#### Using Docker (Recommended)

```bash
# Start MongoDB
docker run -d \
  --name mongodb \
  -p 27017:27017 \
  -e MONGO_INITDB_DATABASE=masked11 \
  mongo:6.0

# Start PostgreSQL
docker run -d \
  --name postgres \
  -p 5432:5432 \
  -e POSTGRES_DB=masked11 \
  -e POSTGRES_USER=masked11_user \
  -e POSTGRES_PASSWORD=masked11_password \
  postgres:15

# Start Redis
docker run -d \
  --name redis \
  -p 6379:6379 \
  redis:7.0-alpine
```

#### Using Local Installation

If you prefer to install databases locally:

**MongoDB:**
```bash
# macOS (using Homebrew)
brew install mongodb-community
brew services start mongodb-community

# Ubuntu/Debian
sudo apt-get install mongodb
sudo systemctl start mongodb
```

**PostgreSQL:**
```bash
# macOS (using Homebrew)
brew install postgresql
brew services start postgresql

# Ubuntu/Debian
sudo apt-get install postgresql postgresql-contrib
sudo systemctl start postgresql
```

**Redis:**
```bash
# macOS (using Homebrew)
brew install redis
brew services start redis

# Ubuntu/Debian
sudo apt-get install redis-server
sudo systemctl start redis-server
```

### Step 3: Install Dependencies

#### Backend Dependencies
```bash
cd backend
go mod download
go mod tidy
cd ..
```

#### Frontend Dependencies
```bash
cd frontend
npm install
cd ..
```

### Step 4: Run Database Migrations

```bash
cd backend
go run scripts/migrate.go
cd ..
```

### Step 5: Start the Servers

#### Start Backend Server
```bash
cd backend
go run cmd/server/main.go
```

You should see output like:
```
🚀 Server starting on port 8080
📊 Health check: http://localhost:8080/health
📈 Metrics: http://localhost:8080/metrics
✅ All databases initialized successfully
```

#### Start Frontend Server (in a new terminal)
```bash
cd frontend
npm run dev
```

You should see output like:
```
- ready started server on 0.0.0.0:3000, url: http://localhost:3000
```

## 🌐 Access Your Application

Once everything is running, you can access:

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Health Check**: http://localhost:8080/health
- **API Metrics**: http://localhost:8080/metrics

## 👤 Default Admin Account

- **Email**: admin@masked11.com
- **Password**: admin123

## 🧪 Test the API

You can test the API endpoints using curl:

```bash
# Health check
curl http://localhost:8080/health

# Get products
curl http://localhost:8080/api/products

# Get categories
curl http://localhost:8080/api/products/categories

# Search products
curl "http://localhost:8080/api/products/search?q=laptop"
```

## 🛠️ Development Commands

### Backend Commands
```bash
cd backend

# Run in development mode
go run cmd/server/main.go

# Run tests
go test ./...

# Run with coverage
go test -cover ./...

# Build for production
go build -o bin/server cmd/server/main.go

# Run migrations
go run scripts/migrate.go
```

### Frontend Commands
```bash
cd frontend

# Start development server
npm run dev

# Build for production
npm run build

# Start production server
npm start

# Run tests
npm test

# Lint code
npm run lint
```

### Database Commands
```bash
# View MongoDB logs
docker logs mongodb

# View PostgreSQL logs
docker logs postgres

# View Redis logs
docker logs redis

# Access MongoDB shell
docker exec -it mongodb mongosh

# Access PostgreSQL shell
docker exec -it postgres psql -U masked11_user -d masked11
```

## 🔧 Troubleshooting

### Common Issues

#### 1. Port Already in Use
```bash
# Check what's using the port
lsof -i :8080
lsof -i :3000

# Kill the process
kill -9 <PID>
```

#### 2. Database Connection Issues
```bash
# Check if databases are running
docker ps

# Restart databases
docker restart mongodb postgres redis
```

#### 3. Go Module Issues
```bash
cd backend
go mod tidy
go mod download
```

#### 4. Node.js Issues
```bash
cd frontend
rm -rf node_modules package-lock.json
npm install
```

#### 5. Permission Issues
```bash
# Make setup script executable
chmod +x setup-localhost.sh

# Fix Docker permissions (if needed)
sudo usermod -aG docker $USER
```

### Health Checks

#### Backend Health
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z",
  "version": "1.0.0",
  "uptime": "1m30s"
}
```

#### Frontend Health
```bash
curl http://localhost:3000
```

Should return the HTML page.

#### Database Health
```bash
# MongoDB
docker exec mongodb mongosh --eval "db.runCommand('ping')"

# PostgreSQL
docker exec postgres pg_isready -U masked11_user

# Redis
docker exec redis redis-cli ping
```

## 📊 Monitoring

### Application Metrics
- **Backend Metrics**: http://localhost:8080/metrics
- **Health Check**: http://localhost:8080/health

### Database Monitoring
```bash
# MongoDB stats
docker exec mongodb mongosh --eval "db.stats()"

# PostgreSQL stats
docker exec postgres psql -U masked11_user -d masked11 -c "SELECT version();"

# Redis info
docker exec redis redis-cli info
```

## 🛑 Stopping Services

### Stop All Services
```bash
# Stop backend
pkill -f "go run"

# Stop frontend
pkill -f "npm run dev"

# Stop databases
docker stop mongodb postgres redis
```

### Clean Up
```bash
# Remove containers
docker rm mongodb postgres redis

# Remove images (optional)
docker rmi mongo:6.0 postgres:15 redis:7.0-alpine

# Clean Docker system
docker system prune -f
```

## 🚀 Next Steps

Once your localhost setup is running:

1. **Explore the Frontend**: Visit http://localhost:3000
2. **Test the API**: Use the health check and product endpoints
3. **Create Products**: Use the admin interface to add products
4. **Test Features**: Try the cart, search, and filtering features
5. **Monitor Performance**: Check the metrics endpoint
6. **Develop**: Start building new features!

## 📚 Additional Resources

- **Backend Documentation**: See `backend/README.md`
- **Frontend Documentation**: See `frontend/README.md`
- **API Documentation**: Check the health and metrics endpoints
- **Architecture**: See `docs/architecture/system-design.md`

---

**Happy coding! 🎉** 