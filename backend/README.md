# Masked 11 Ecommerce Backend

A high-performance, production-ready ecommerce API built with Go, Fiber, and modern best practices.

## 🚀 Features

### Core Ecommerce Features
- **Authentication & Authorization**: JWT-based auth with role-based access control
- **Product Management**: CRUD operations, search, filtering, pagination, categories
- **Shopping Cart**: Redis-based cart with session management
- **Order Management**: Order creation, tracking, status updates, cancellation
- **User Management**: Registration, login, profile management, admin roles

### Security & Performance
- **Security Headers**: Helmet middleware with XSS, CSP, HSTS protection
- **Rate Limiting**: Adaptive rate limiting with Redis storage and burst protection
- **CORS**: Configurable cross-origin resource sharing
- **Request Validation**: Comprehensive input validation with detailed error messages
- **Compression**: Response compression for better performance
- **Caching**: Redis-based caching with TTL and invalidation

### Monitoring & Observability
- **Metrics Collection**: Request/response tracking, error monitoring, performance metrics
- **Health Checks**: Comprehensive health check endpoints
- **Prometheus Integration**: Metrics export in Prometheus format
- **Structured Logging**: Request ID tracking, performance logging
- **Database Monitoring**: Connection pool monitoring, query tracking

### Production Features
- **Graceful Shutdown**: Proper cleanup on server termination
- **Connection Pooling**: Optimized database connections
- **Error Handling**: Standardized error responses with request tracking
- **Environment Configuration**: Comprehensive environment variable management
- **Docker Support**: Multi-stage Dockerfile for production builds

## 📋 Prerequisites

- Go 1.24.4+
- Docker & Docker Compose
- MongoDB 6.0+
- PostgreSQL 15+
- Redis 7.0+

## 🛠️ Installation & Setup

### 1. Clone and Setup
```bash
git clone <repository-url>
cd backend
```

### 2. Environment Configuration
```bash
# Copy environment template
cp .env.example .env

# Edit environment variables
nano .env
```

### 3. Database Setup
```bash
# Start databases with Docker
docker-compose up -d mongodb postgres redis

# Run migrations
go run scripts/migrate.go
```

### 4. Install Dependencies
```bash
go mod download
```

### 5. Run Development Server
```bash
# Using Makefile
make dev

# Or directly
go run cmd/server/main.go
```

## 🔧 Configuration

### Environment Variables

#### Server Configuration
```env
APP_PORT=8080
APP_ENV=development
APP_NAME=Masked11-API
APP_VERSION=1.0.0
```

#### Database Configuration
```env
# MongoDB
MONGO_URI=mongodb://localhost:27017
MONGO_DB=masked11

# PostgreSQL
POSTGRES_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres

# Redis
REDIS_URI=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

#### Security Configuration
```env
SESSION_SECRET=your-super-secret-key-change-in-production
JWT_SECRET=your-jwt-secret-key-change-in-production
JWT_EXPIRY=168h
BCRYPT_COST=12

# CORS
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
ALLOWED_HEADERS=Origin,Content-Type,Accept,Authorization,X-Requested-With
```

#### Rate Limiting
```env
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=1m
```

#### Monitoring
```env
METRICS_ENABLED=true
METRICS_PORT=9090
HEALTH_CHECK_ENABLED=true
```

## 📚 API Documentation

### Authentication Endpoints

#### Register User
```http
POST /api/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePass123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": "507f1f77bcf86cd799439011",
      "email": "john@example.com",
      "name": "John Doe",
      "role": "user"
    }
  },
  "timestamp": "2024-01-01T12:00:00Z",
  "request_id": "req-123456"
}
```

#### Login User
```http
POST /api/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "SecurePass123"
}
```

### Product Endpoints

#### Get All Products
```http
GET /api/products?page=1&limit=12&search=phone&category=electronics&minPrice=100&maxPrice=1000&sortBy=price&sortOrder=asc
```

#### Get Product by ID
```http
GET /api/products/{id}
```

#### Search Products
```http
GET /api/products/search?q=iphone&page=1&limit=12
```

### Cart Endpoints

#### Add to Cart
```http
POST /api/cart/add
Content-Type: application/json

{
  "productID": "507f1f77bcf86cd799439011",
  "quantity": 2,
  "size": "M",
  "price": 99.99
}
```

#### Get Cart
```http
GET /api/cart
```

### Order Endpoints

#### Create Order
```http
POST /api/orders
Authorization: Bearer {token}
Content-Type: application/json

{
  "items": [
    {
      "productID": "507f1f77bcf86cd799439011",
      "quantity": 2,
      "price": 99.99
    }
  ],
  "shippingAddress": "123 Main St, City, State 12345"
}
```

### Admin Endpoints

#### Create Product (Admin Only)
```http
POST /api/admin/products
Authorization: Bearer {admin_token}
Content-Type: application/json

{
  "name": "iPhone 15 Pro",
  "description": "Latest iPhone with advanced features",
  "price": 999.99,
  "category": "electronics",
  "stock": 50,
  "images": ["image1.jpg", "image2.jpg"]
}
```

## 🔍 Monitoring Endpoints

### Health Check
```http
GET /health
```

### Metrics
```http
GET /metrics
Authorization: Bearer {admin_token}
```

### Prometheus Metrics
```http
GET /metrics/prometheus
Authorization: Bearer {admin_token}
```

### Database Metrics
```http
GET /metrics/database
Authorization: Bearer {admin_token}
```

## 🧪 Testing

### Run All Tests
```bash
make test
```

### Run Tests with Coverage
```bash
make test-coverage
```

### Run Specific Test
```bash
go test ./internal/utils -v
```

## 🐳 Docker

### Build Image
```bash
docker build -t masked11-backend .
```

### Run with Docker Compose
```bash
docker-compose up -d
```

### Production Build
```bash
docker build -f Dockerfile.prod -t masked11-backend:prod .
```

## 📊 Performance Features

### Rate Limiting
- **Default**: 100 requests per minute per IP/user
- **Auth Endpoints**: 5 requests per 15 minutes
- **Admin Endpoints**: 1000 requests per minute
- **Adaptive**: Adjusts based on server load

### Caching
- **Product Cache**: 1 hour TTL
- **Category Cache**: 24 hours TTL
- **User Session**: 7 days TTL
- **Cart Data**: Session-based with Redis

### Database Optimization
- **Connection Pooling**: Configurable pool sizes
- **Query Optimization**: Indexed queries
- **Connection Monitoring**: Health checks and metrics

## 🔒 Security Features

### Input Validation
- **Email Validation**: RFC 5322 compliant
- **Password Strength**: Minimum 8 chars, uppercase, lowercase, number
- **Phone Validation**: International format support
- **Data Sanitization**: XSS protection

### Security Headers
- **XSS Protection**: `X-XSS-Protection: 1; mode=block`
- **Content Security Policy**: Restrictive CSP
- **HSTS**: HTTPS enforcement
- **Frame Options**: Clickjacking protection

### Authentication
- **JWT Tokens**: Secure token-based auth
- **Session Management**: Redis-based sessions
- **Password Hashing**: Bcrypt with configurable cost
- **Role-Based Access**: Admin and user roles

## 🚀 Deployment

### Production Checklist
- [ ] Update environment variables for production
- [ ] Set secure JWT and session secrets
- [ ] Configure CORS for production domains
- [ ] Set up SSL/TLS certificates
- [ ] Configure database backups
- [ ] Set up monitoring and alerting
- [ ] Configure rate limiting for production load
- [ ] Set up logging aggregation

### Environment Variables for Production
```env
APP_ENV=production
APP_PORT=8080

# Use strong secrets
JWT_SECRET=your-very-long-and-secure-jwt-secret
SESSION_SECRET=your-very-long-and-secure-session-secret

# Production databases
MONGO_URI=mongodb://user:pass@prod-mongo:27017/masked11
POSTGRES_URL=postgres://user:pass@prod-postgres:5432/masked11?sslmode=require
REDIS_URI=prod-redis:6379

# Production CORS
ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com

# Production rate limiting
RATE_LIMIT_REQUESTS=200
RATE_LIMIT_WINDOW=1m
```

## 📈 Monitoring & Observability

### Metrics Available
- **Request Count**: Total requests processed
- **Error Count**: Total errors encountered
- **Response Time**: Average response time
- **Rate Limit Exceeded**: Rate limiting violations
- **Database Queries**: Query count and errors
- **Cache Performance**: Hit/miss rates
- **Active Connections**: Current connection count

### Health Checks
- **Database Connectivity**: All databases
- **Redis Connectivity**: Cache availability
- **Service Status**: Overall service health
- **Dependencies**: External service health

### Logging
- **Structured Logs**: JSON format with request ID
- **Performance Logging**: Response time tracking
- **Error Logging**: Detailed error information
- **Security Logging**: Authentication events

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Check the documentation
- Review the API examples

---

**Built with ❤️ using Go, Fiber, and modern best practices**
