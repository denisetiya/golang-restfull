# Setup dan Testing Guide

## Persiapan Environment

### 1. Install Dependencies
```bash
go mod tidy
```

### 2. Setup MongoDB

#### Option A: Menggunakan Docker (Recommended)
```bash
# Start MongoDB dengan Docker Compose
docker-compose up -d mongodb

# Cek status MongoDB
docker-compose ps
```

#### Option B: Install MongoDB Local
- Download dan install MongoDB dari [https://www.mongodb.com/try/download/community](https://www.mongodb.com/try/download/community)
- Start MongoDB service

### 3. Konfigurasi Environment
Copy dan edit file `.env`:
```bash
cp .env.example .env
```

Edit `.env` dengan konfigurasi yang sesuai:
```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=rest_api_db
JWT_SECRET=your-very-secret-key-here
JWT_EXPIRE_HOURS=24
SERVER_PORT=8080
```

## Menjalankan Aplikasi

### Option 1: Quick Start
```bash
chmod +x start.sh
./start.sh
```

### Option 2: Manual
```bash
# Build aplikasi
go build -o bin/rest-api main.go

# Jalankan
./bin/rest-api
```

### Option 3: Development Mode
```bash
go run main.go
```

## Testing API

### 1. Health Check
```bash
curl http://localhost:8080/health
```

### 2. Register User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 3. Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

Response akan berisi token JWT yang perlu disimpan untuk request selanjutnya.

### 4. Create Product (Requires Authentication)
```bash
# Ganti YOUR_JWT_TOKEN dengan token dari response login
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

### 5. Get All Products
```bash
curl http://localhost:8080/api/v1/products
```

### 6. Get Product by ID
```bash
# Ganti PRODUCT_ID dengan ID produk yang valid
curl http://localhost:8080/api/v1/products/PRODUCT_ID
```

## Troubleshooting

### MongoDB Connection Issues
1. Pastikan MongoDB running di port 27017
2. Cek konfigurasi MONGO_URI di file .env
3. Test koneksi MongoDB:
   ```bash
   # Jika MongoDB local
   mongosh
   
   # Jika menggunakan Docker
   docker exec -it golang_mongodb_1 mongosh
   ```

### Build Errors
1. Pastikan Go versi 1.19+
2. Run `go mod tidy` untuk update dependencies
3. Cek error messages untuk dependency issues

### Permission Errors (Docker)
```bash
# Add user to docker group (Linux)
sudo usermod -aG docker $USER
# Logout and login again

# Or run with sudo
sudo docker-compose up -d mongodb
```

## Struktur Database MongoDB

### Collections:
- `users` - Data pengguna
- `products` - Data produk

### Sample Data Structure:

#### User Document:
```json
{
  "_id": ObjectId("..."),
  "name": "John Doe",
  "email": "john@example.com",
  "password": "hashed_password",
  "created_at": ISODate("..."),
  "updated_at": ISODate("...")
}
```

#### Product Document:
```json
{
  "_id": ObjectId("..."),
  "name": "Laptop Gaming",
  "description": "Laptop gaming dengan spesifikasi tinggi",
  "price": 15000000,
  "stock": 10,
  "user_id": ObjectId("..."),
  "created_at": ISODate("..."),
  "updated_at": ISODate("...")
}
```

## Best Practices Implemented

1. **Validation**: Input validation menggunakan struct tags
2. **Security**: Password hashing dengan bcrypt, JWT authentication
3. **Error Handling**: Consistent error responses
4. **Clean Code**: Separation of concerns dengan layer architecture
5. **Configuration**: Environment-based configuration
6. **Documentation**: Comprehensive API documentation
7. **Testing**: Structure ready for unit/integration tests
