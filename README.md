# Masked 11 - Premium Ecommerce Platform

A full-fledged, high-performance ecommerce platform built with modern technologies for fast loading and exceptional user experience.

## 🚀 Features

### Core Ecommerce Features
- **Product Management**: Full CRUD operations with categories, sizes, and inventory tracking
- **Advanced Search & Filtering**: Real-time search with price range, category, and rating filters
- **Shopping Cart**: Redis-based cart with session management
- **Order Management**: Complete order lifecycle with status tracking
- **User Authentication**: JWT-based auth with role-based access control
- **Admin Dashboard**: Comprehensive admin interface for product and order management

### Performance & Scalability
- **Multi-Database Architecture**: MongoDB for products/users, PostgreSQL for orders, Redis for sessions/cart
- **Caching Strategy**: Redis caching for improved response times
- **Pagination**: Efficient pagination for large product catalogs
- **Image Optimization**: Optimized image handling and CDN ready
- **API Rate Limiting**: Built-in rate limiting for API protection

### Frontend Features
- **Modern UI/UX**: Beautiful, responsive design with Tailwind CSS
- **Fast Loading**: Optimized with Next.js 15 and App Router
- **Mobile-First**: Fully responsive design for all devices
- **Search & Filter**: Real-time search with advanced filtering options
- **Shopping Cart**: Persistent cart with quantity management

## 🛠 Tech Stack

### Backend
- **Go 1.24.4** with Fiber framework
- **MongoDB** - Product and user data
- **PostgreSQL** - Order management
- **Redis** - Session and cart storage
- **JWT** - Authentication
- **Docker** - Containerization

### Frontend
- **Next.js 15** with App Router
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **React 19** - UI framework

### Infrastructure
- **Docker Compose** - Development environment
- **Nginx** - Reverse proxy
- **Caddy** - Web server
- **Prometheus** - Monitoring

## 📁 Project Structure

```
Masked11/
├── backend/                 # Go API server
│   ├── cmd/server/         # Main application
│   ├── internal/           # Business logic
│   │   ├── handlers/      # HTTP handlers
│   │   ├── services/      # Business services
│   │   ├── repositories/  # Data access layer
│   │   ├── models/        # Data models
│   │   ├── middleware/    # HTTP middleware
│   │   └── database/      # Database connections
│   └── migrations/        # Database migrations
├── frontend/              # Next.js application
│   ├── src/app/          # App Router pages
│   │   ├── (shop)/       # Shop pages
│   │   ├── (admin)/      # Admin pages
│   │   └── (auth)/       # Authentication pages
│   └── public/           # Static assets
├── infrastructure/        # Deployment configs
│   ├── docker/           # Docker configurations
│   ├── nginx/            # Nginx configs
│   └── monitoring/       # Monitoring setup
└── docs/                # Documentation
```

## 🚀 Quick Start

### Prerequisites
- Go 1.24.4+
- Node.js 18+
- Docker & Docker Compose
- MongoDB, PostgreSQL, Redis

### Development Setup

1. **Clone the repository**
```bash
git clone https://github.com/your-username/Masked11.git
cd Masked11
```

2. **Set up environment variables**
```bash
# Backend
cp backend/.env.example backend/.env
# Edit backend/.env with your database credentials

# Frontend
cp frontend/.env.example frontend/.env
```

3. **Start with Docker Compose**
```bash
docker-compose -f infrastructure/docker/docker-compose.yml up -d
```

4. **Run migrations**
```bash
cd backend
go run scripts/migrate.go
```

5. **Start the backend**
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

6. **Start the frontend**
```bash
cd frontend
npm install
npm run dev
```

### Production Deployment

1. **Build the application**
```bash
# Backend
cd backend
go build -o bin/server cmd/server/main.go

# Frontend
cd frontend
npm run build
```

2. **Deploy with Docker**
```bash
docker-compose -f infrastructure/docker/docker-compose.prod.yml up -d
```

## 📚 API Documentation

### Authentication Endpoints
- `POST /api/register` - User registration
- `POST /api/login` - User login
- `POST /api/logout` - User logout

### Product Endpoints
- `GET /api/products` - Get all products with filtering
- `GET /api/products/search` - Search products
- `GET /api/products/categories` - Get product categories
- `GET /api/products/featured` - Get featured products
- `GET /api/products/:id` - Get product by ID

### Cart Endpoints
- `POST /api/cart/add` - Add item to cart
- `GET /api/cart` - Get cart contents
- `DELETE /api/cart/remove/:id` - Remove item from cart
- `DELETE /api/cart/clear` - Clear cart

### Order Endpoints
- `POST /api/orders` - Create order
- `GET /api/orders` - Get user orders
- `GET /api/orders/:id` - Get order by ID
- `PUT /api/orders/:id/cancel` - Cancel order

## 🔧 Configuration

### Environment Variables

**Backend (.env)**
```env
# Database
MONGO_URI=mongodb://localhost:27017
MONGO_DB=masked11
POSTGRES_URL=postgres://user:pass@localhost:5432/masked11
REDIS_URI=localhost:6379

# JWT
SESSION_SECRET=your-secret-key

# Admin
ADMIN_EMAIL=admin@masked11.com
ADMIN_PASSWORD=admin123
ADMIN_NAME=Admin

# Server
APP_PORT=8080
```

**Frontend (.env)**
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_SITE_URL=http://localhost:3000
```

## 🧪 Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm test
```

## 📊 Performance Optimizations

### Backend
- **Database Indexing**: Optimized indexes for fast queries
- **Connection Pooling**: Efficient database connections
- **Caching**: Redis caching for frequently accessed data
- **Pagination**: Efficient pagination for large datasets

### Frontend
- **Image Optimization**: Next.js Image component
- **Code Splitting**: Automatic code splitting
- **Static Generation**: Pre-rendered pages where possible
- **CDN Ready**: Optimized for CDN deployment

## 🔒 Security Features

- **JWT Authentication**: Secure token-based authentication
- **Role-Based Access**: Admin and user role management
- **Input Validation**: Comprehensive input validation
- **CORS Protection**: Proper CORS configuration
- **Rate Limiting**: API rate limiting protection

## 📈 Monitoring & Analytics

- **Health Checks**: `/health` endpoint for monitoring
- **Prometheus Metrics**: Built-in metrics collection
- **Error Logging**: Comprehensive error logging
- **Performance Monitoring**: Response time tracking

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support, email support@masked11.com or create an issue in the repository.

## 🗺 Roadmap

- [ ] Payment gateway integration (Stripe/PayPal)
- [ ] Email notifications
- [ ] Product reviews and ratings
- [ ] Wishlist functionality
- [ ] Advanced analytics dashboard
- [ ] Mobile app (React Native)
- [ ] Multi-language support
- [ ] Advanced inventory management
- [ ] Affiliate program
- [ ] Loyalty points system

---

**Masked 11** - Premium Ecommerce Platform 🛍️
