package models

// Request DTOs
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"omitempty,min=2,max=100"`
	Email string `json:"email" validate:"omitempty,email"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"max=500"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"omitempty,min=2,max=100"`
	Description string  `json:"description" validate:"omitempty,max=500"`
	Price       float64 `json:"price" validate:"omitempty,gt=0"`
	Stock       int     `json:"stock" validate:"omitempty,gte=0"`
}

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items" validate:"required,min=1,dive"`
}

type OrderItemRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,gt=0"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending processing completed cancelled"`
}

// Response DTOs
type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type ProductResponse struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float64      `json:"price"`
	Stock       int          `json:"stock"`
	User        UserResponse `json:"user"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

type OrderResponse struct {
	ID          string              `json:"id"`
	User        UserResponse        `json:"user"`
	TotalAmount float64             `json:"total_amount"`
	Status      string              `json:"status"`
	OrderItems  []OrderItemResponse `json:"order_items"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
}

type OrderItemResponse struct {
	ID       string          `json:"id"`
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
	Price    float64         `json:"price"`
}

// Generic responses
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int64       `json:"total"`
}
