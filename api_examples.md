# REST API Testing Examples

## Health Check
```bash
curl -X GET http://localhost:8080/health
```

## User Registration
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

## User Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

## Get User Profile (requires token)
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Update User Profile (requires token)
```bash
curl -X PUT http://localhost:8080/api/v1/users/profile \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "John Updated"
  }'
```

## Get All Products
```bash
curl -X GET "http://localhost:8080/api/v1/products?page=1&limit=10"
```

## Create Product (requires token)
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Sample Product",
    "description": "A sample product for testing",
    "price": 29.99,
    "stock": 100
  }'
```

## Get Product by ID
```bash
curl -X GET http://localhost:8080/api/v1/products/1
```

## Get My Products (requires token)
```bash
curl -X GET "http://localhost:8080/api/v1/products/my?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Update Product (requires token)
```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Updated Product Name",
    "price": 39.99,
    "stock": 50
  }'
```

## Delete Product (requires token)
```bash
curl -X DELETE http://localhost:8080/api/v1/products/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Get All Users (requires token)
```bash
curl -X GET "http://localhost:8080/api/v1/users?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Delete User Account (requires token)
```bash
curl -X DELETE http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Example Response Format

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Operation failed",
  "error": "Detailed error message"
}
```

### Paginated Response
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [...],
  "page": 1,
  "limit": 10,
  "total": 100
}
```
