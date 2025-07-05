README.md

# RESTful API with Go, Gin, MongoDB, and JWT

A comprehensive RESTful API built with Go using best practices including validation, MongoDB, JWT authentication, and clean architecture.

## Features

- **Clean Architecture**: Separated into layers (handlers, services, repositories)
- **JWT Authentication**: Secure token-based authentication
- **Input Validation**: Request validation using go-playground/validator
- **MongoDB**: Database operations using MongoDB driver
- **Pagination**: Support for paginated responses
- **CORS**: Cross-origin resource sharing support
- **Error Handling**: Comprehensive error handling middleware
- **MongoDB**: Database support with automatic connection

## Project Structure

```
├── cmd/
│   └── main.go                 # Alternative main entry point
├── internal/
│   ├── config/
│   │   ├── config.go           # Configuration management
│   │   └── database.go         # Database connection
│   ├── handlers/
│   │   ├── user_handler.go     # User HTTP handlers
│   │   └── product_handler.go  # Product HTTP handlers
│   ├── middleware/
│   │   └── middleware.go       # JWT auth, CORS, error handling
│   ├── models/
│   │   ├── models.go           # Database models
│   │   └── dto.go              # Request/Response DTOs
│   ├── repositories/
│   │   ├── user_repository.go  # User database operations
│   │   └── product_repository.go # Product database operations
│   ├── services/
│   │   ├── user_service.go     # User business logic
│   │   └── product_service.go  # Product business logic
│   └── utils/
│       └── utils.go            # Utility functions (JWT, password hashing)
├── .env                        # Environment variables
├── go.mod                      # Go module file
└── main.go                     # Main entry point
```

## Setup Instructions

### Prerequisites

- Go 1.19 or higher
- MongoDB database
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd golang
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup MongoDB database**
   - Create a MongoDB database
   - Update the `.env` file with your database credentials

4. **Configure environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

5. **Run the application**
   ```bash
   # Option 1: Direct run (requires MongoDB running locally)
   go run main.go
   
   # Option 2: Build and run
   go build -o bin/rest-api main.go
   ./bin/rest-api
   
   # Option 3: Use start script (with Docker MongoDB)
   ./start.sh
   ```

The server will start on port 8080 by default.

## Quick Start with Docker MongoDB

If you have Docker installed, you can quickly start the application:

```bash
# Make scripts executable
chmod +x start.sh build.sh

# Start MongoDB and application
./start.sh
```

## Manual Setup

If you prefer to setup MongoDB manually:

1. **Install MongoDB** on your system
2. **Start MongoDB service**
3. **Run the application**:
   ```bash
   go run main.go
   ```

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user

### Users (Protected)

- `GET /api/v1/users/profile` - Get user profile
- `PUT /api/v1/users/profile` - Update user profile
- `DELETE /api/v1/users/profile` - Delete user account
- `GET /api/v1/users/` - Get all users (with pagination)

### Products

- `GET /api/v1/products/` - Get all products (with pagination)
- `GET /api/v1/products/:id` - Get product by ID

### Products (Protected)

- `POST /api/v1/products/` - Create new product
- `GET /api/v1/products/my` - Get user's products
- `PUT /api/v1/products/:id` - Update product
- `DELETE /api/v1/products/:id` - Delete product

### Health Check

- `GET /health` - Health check endpoint

## Example API Usage

### Register User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create Product (requires authentication)

```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "name": "Sample Product",
    "description": "A sample product",
    "price": 29.99,
    "stock": 100
  }'
```

## Environment Variables

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=rest_api_db

JWT_SECRET=your-secret-key-here
JWT_EXPIRE_HOURS=24

SERVER_PORT=8080
```

## Database Models

### User
- ID (ObjectID)
- Name (required, 2-100 chars)
- Email (required, unique, valid email)
- Password (required, min 6 chars, hashed)
- CreatedAt, UpdatedAt

### Product
- ID (ObjectID)
- Name (required, 2-100 chars)
- Description (optional, max 500 chars)
- Price (required, > 0)
- Stock (required, >= 0)
- UserID (ObjectID reference)
- CreatedAt, UpdatedAt

## Best Practices Implemented

1. **Clean Architecture**: Separation of concerns with clear layer boundaries
2. **Validation**: Input validation using struct tags and go-playground/validator
3. **Security**: JWT authentication, password hashing with bcrypt
4. **Error Handling**: Comprehensive error handling with proper HTTP status codes
5. **Database**: MongoDB with proper indexing and relationships
6. **Configuration**: Environment-based configuration
7. **Middleware**: CORS, authentication, and error handling middleware
8. **Pagination**: Proper pagination for list endpoints
9. **Logging**: Structured logging for debugging and monitoring
10. **Code Organization**: Clear folder structure and naming conventions

## Testing

Run tests (when implemented):
```bash
go test ./...
```

## Production Considerations

1. **Database**: Use connection pooling and proper indexing
2. **Caching**: Implement Redis for caching frequently accessed data
3. **Rate Limiting**: Add rate limiting middleware
4. **Monitoring**: Add metrics and health checks
5. **Logging**: Structured logging with proper log levels
6. **Security**: Input sanitization, SQL injection prevention
7. **Documentation**: API documentation with Swagger/OpenAPI
8. **Testing**: Unit tests, integration tests, and end-to-end tests


## License

This project is licensed under the MIT License.
