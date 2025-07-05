# 🚀 RESTful API dengan MongoDB - Go Gin Framework

Selamat! Anda telah berhasil membuat RESTful API yang lengkap dengan teknologi modern dan best practices.

## ✅ Apa yang Telah Dibuat

### 🏗️ Arsitektur Clean Code
- **Handlers**: Menangani HTTP requests/responses
- **Services**: Business logic dan validasi
- **Repositories**: Database operations
- **Models**: Data structures dan DTOs
- **Middleware**: Authentication, CORS, error handling
- **Utils**: Utility functions (JWT, password hashing)

### 🛠️ Teknologi yang Digunakan
- **Go**: Programming language
- **Gin**: HTTP web framework
- **MongoDB**: NoSQL database
- **JWT**: Token-based authentication
- **Validator**: Input validation
- **Bcrypt**: Password hashing
- **Docker**: Containerization untuk MongoDB

### 📁 Struktur Proyek
```
golang/
├── cmd/main.go                 # Alternative entry point
├── internal/
│   ├── config/                 # Konfigurasi dan database
│   ├── handlers/               # HTTP handlers
│   ├── middleware/             # Middleware (auth, cors)
│   ├── models/                 # Data models & DTOs
│   ├── repositories/           # Database operations
│   ├── services/               # Business logic
│   └── utils/                  # Utility functions
├── .env                        # Environment variables
├── docker-compose.yml          # MongoDB container
├── start.sh                    # Auto-start script
├── README.md                   # Documentation
├── SETUP.md                    # Setup guide
└── main.go                     # Main entry point
```

## 🌟 Fitur yang Tersedia

### 🔐 Authentication & Authorization
- User registration dengan validasi
- Login dengan JWT token
- Password hashing dengan bcrypt
- Protected routes dengan middleware

### 👤 User Management
- ✅ `POST /api/v1/auth/register` - Daftar user baru
- ✅ `POST /api/v1/auth/login` - Login user
- ✅ `GET /api/v1/users/profile` - Lihat profil (protected)
- ✅ `PUT /api/v1/users/profile` - Update profil (protected)
- ✅ `DELETE /api/v1/users/profile` - Hapus akun (protected)
- ✅ `GET /api/v1/users/` - List semua users (protected)

### 📦 Product Management
- ✅ `GET /api/v1/products/` - List semua produk (public + pagination)
- ✅ `GET /api/v1/products/:id` - Detail produk (public)
- ✅ `POST /api/v1/products/` - Buat produk baru (protected)
- ✅ `GET /api/v1/products/my` - Produk saya (protected)
- ✅ `PUT /api/v1/products/:id` - Update produk (protected, owner only)
- ✅ `DELETE /api/v1/products/:id` - Hapus produk (protected, owner only)

### 🔧 Fitur Tambahan
- ✅ Health check endpoint
- ✅ Pagination untuk list endpoints
- ✅ CORS support
- ✅ Error handling yang konsisten
- ✅ Input validation
- ✅ Environment-based configuration

## 🚀 Cara Menjalankan

### Quick Start (Recommended)
```bash
# Make script executable
chmod +x start.sh

# Start MongoDB dan aplikasi
./start.sh
```

### Manual Setup
```bash
# 1. Install dependencies
go mod tidy

# 2. Start MongoDB (jika menggunakan Docker)
docker-compose up -d mongodb

# 3. Run aplikasi
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## 🧪 Testing API

### 1. Health Check
```bash
curl http://localhost:8080/health
```

### 2. Register User Baru
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com", 
    "password": "password123"
  }'
```

### 3. Login dan Dapatkan Token
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 4. Buat Produk (Gunakan Token dari Login)
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Laptop Gaming",
    "description": "Laptop gaming dengan spesifikasi tinggi",
    "price": 15000000,
    "stock": 10
  }'
```

## 🎯 Best Practices yang Diterapkan

1. **Clean Architecture**: Pemisahan concerns yang jelas
2. **Validation**: Validasi input menggunakan struct tags
3. **Security**: Password hashing, JWT authentication
4. **Error Handling**: Response error yang konsisten
5. **Configuration**: Environment-based config
6. **Database**: MongoDB dengan structure yang baik
7. **Middleware**: CORS, authentication, error handling
8. **Pagination**: Pagination untuk performa yang baik
9. **Documentation**: Dokumentasi yang lengkap
10. **Code Organization**: Folder structure yang rapi

## 📚 Dokumentasi Lengkap

- `README.md` - Overview dan panduan dasar
- `SETUP.md` - Panduan setup detail dan troubleshooting
- `api_examples.md` - Contoh penggunaan API lengkap

## 🔮 Pengembangan Selanjutnya

Untuk mengembangkan API ini lebih lanjut, Anda bisa:

1. **Menambah fitur Order Management** (model sudah ada)
2. **Implementasi testing** (unit test, integration test)
3. **Rate limiting** untuk API protection
4. **Logging** yang lebih terstruktur
5. **Metrics dan monitoring**
6. **API documentation** dengan Swagger
7. **Caching** dengan Redis
8. **File upload** untuk product images
9. **Email verification** untuk user registration
10. **Role-based access control** (admin, user, etc.)

## 🎉 Selamat!

Anda telah berhasil membuat RESTful API yang production-ready dengan:
- ✅ Database MongoDB
- ✅ JWT Authentication  
- ✅ Input Validation
- ✅ Clean Architecture
- ✅ Error Handling
- ✅ Best Practices

API ini siap untuk dikembangkan lebih lanjut sesuai kebutuhan bisnis Anda!
