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

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	user.UpdatedAt = time.Now()

	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": user.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *UserRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func (r *UserRepository) GetAll(offset, limit int) ([]models.User, int64, error) {
	// Get total count
	total, err := r.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Get users with pagination
	opts := options.Find()
	opts.SetSkip(int64(offset))
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.TODO())

	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
