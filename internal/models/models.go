package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required,min=2,max=100"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"-" bson:"password" validate:"required,min=6"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required,min=2,max=100"`
	Description string             `json:"description" bson:"description" validate:"max=500"`
	Price       float64            `json:"price" bson:"price" validate:"required,gt=0"`
	Stock       int                `json:"stock" bson:"stock" validate:"required,gte=0"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	User        User               `json:"user,omitempty" bson:"-"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type Order struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	User        User               `json:"user,omitempty" bson:"-"`
	TotalAmount float64            `json:"total_amount" bson:"total_amount" validate:"required,gt=0"`
	Status      string             `json:"status" bson:"status" validate:"oneof=pending processing completed cancelled"`
	OrderItems  []OrderItem        `json:"order_items" bson:"order_items"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type OrderItem struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Product   Product            `json:"product,omitempty" bson:"-"`
	Quantity  int                `json:"quantity" bson:"quantity" validate:"required,gt=0"`
	Price     float64            `json:"price" bson:"price" validate:"required,gt=0"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
