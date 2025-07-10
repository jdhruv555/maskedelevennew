package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/interfaces"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) CreateUser(user *models.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(id primitive.ObjectID, updated bson.M) error {
	updated["updatedAt"] = time.Now()
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updated})
	return err
}

func (r *userRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
