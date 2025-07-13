# 🚀 Quick Start - Masked 11 Ecommerce

## ⚡ One-Command Setup

Run this single command to set up everything:

```bash
./setup-localhost.sh
```

This will:
- ✅ Start all databases (MongoDB, PostgreSQL, Redis)
- ✅ Install all dependencies
- ✅ Run database migrations
- ✅ Start backend server (port 8080)
- ✅ Start frontend server (port 3000)
- ✅ Open your browser automatically

## 🌐 Access Your Website

Once the setup is complete, open your browser to:

- **🏠 Frontend**: http://localhost:3000
- **🔧 Backend API**: http://localhost:8080
- **🏥 Health Check**: http://localhost:8080/health

## 👤 Admin Login

- **Email**: admin@masked11.com
- **Password**: admin123

## 🧪 Test Your Setup

```bash
# Test backend health
curl http://localhost:8080/health

# Test products API
curl http://localhost:8080/api/products

# Test frontend
curl http://localhost:3000
```

## 🛑 Stop Everything

Press `Ctrl+C` in the terminal where you ran the setup script, or run:

```bash
# Stop all services
pkill -f "go run"
pkill -f "npm run dev"
docker stop mongodb postgres redis
```

## 🔧 Manual Setup (Alternative)

If the automated script doesn't work, follow the detailed guide in `LOCALHOST-SETUP.md`.

## 🆘 Troubleshooting

### Common Issues:

1. **Port already in use**: Kill existing processes
   ```bash
   lsof -i :8080 -i :3000
   kill -9 <PID>
   ```

2. **Docker not running**: Start Docker Desktop

3. **Permission denied**: Make script executable
   ```bash
   chmod +x setup-localhost.sh
   ```

4. **Database connection issues**: Restart databases
   ```bash
   docker restart mongodb postgres redis
   ```

## 📊 What You'll See

### Frontend (http://localhost:3000)
- Modern ecommerce homepage
- Product catalog with filtering
- Shopping cart functionality
- User authentication
- Admin dashboard

### Backend (http://localhost:8080)
- RESTful API endpoints
- Real-time metrics
- Health monitoring
- Performance optimization
- Caching system

### Features Available:
- ✅ Product browsing and search
- ✅ Shopping cart management
- ✅ User registration/login
- ✅ Admin product management
- ✅ Real-time performance monitoring
- ✅ Advanced caching
- ✅ Rate limiting
- ✅ Security features

## 🎉 You're Ready!

Your Masked 11 ecommerce website is now running locally with:
- **High-performance backend** with Go
- **Modern frontend** with Next.js
- **Professional caching** with Redis
- **Scalable databases** with MongoDB & PostgreSQL
- **Real-time monitoring** and metrics

Start exploring and building! 🚀 