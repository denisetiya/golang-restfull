package repositories

import (
	"context"
	"errors"
	"time"

	"rest-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
	userRepo   *UserRepository
}

func NewProductRepository(db *mongo.Database, userRepo *UserRepository) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
		userRepo:   userRepo,
	}
}

func (r *ProductRepository) Create(product *models.Product) error {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(context.TODO(), product)
	return err
}

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product models.Product
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Load user data
	if !product.UserID.IsZero() {
		user, err := r.userRepo.GetByID(product.UserID.Hex())
		if err == nil {
			product.User = *user
		}
	}

	return &product, nil
}

func (r *ProductRepository) GetAll(offset, limit int) ([]models.Product, int64, error) {
	// Get total count
	total, err := r.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Get products with pagination
	opts := options.Find()
	opts.SetSkip(int64(offset))
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.TODO())

	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, 0, err
	}

	// Load user data for each product
	for i := range products {
		if !products[i].UserID.IsZero() {
			user, err := r.userRepo.GetByID(products[i].UserID.Hex())
			if err == nil {
				products[i].User = *user
			}
		}
	}

	return products, total, nil
}

func (r *ProductRepository) GetByUserID(userID string, offset, limit int) ([]models.Product, int64, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, 0, err
	}

	filter := bson.M{"user_id": objID}

	// Get total count
	total, err := r.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, 0, err
	}

	// Get products with pagination
	opts := options.Find()
	opts.SetSkip(int64(offset))
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.TODO())

	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, 0, err
	}

	// Load user data for each product
	for i := range products {
		if !products[i].UserID.IsZero() {
			user, err := r.userRepo.GetByID(products[i].UserID.Hex())
			if err == nil {
				products[i].User = *user
			}
		}
	}

	return products, total, nil
}

func (r *ProductRepository) Update(product *models.Product) error {
	product.UpdatedAt = time.Now()

	filter := bson.M{"_id": product.ID}
	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"updated_at":  product.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *ProductRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func (r *ProductRepository) UpdateStock(id string, stock int) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"stock":      stock,
			"updated_at": time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
