package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Shrey-Yash/Masked11/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id primitive.ObjectID) (*models.User, error)
	UpdateUser(id primitive.ObjectID, updated bson.M) error
	DeleteUser(id primitive.ObjectID) error
}
