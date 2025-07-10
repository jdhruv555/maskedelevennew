package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name      string               `bson:"name" json:"name" validate:"required"`
	Email     string               `bson:"email" json:"email" validate:"required,email"`
	Password  string               `bson:"password" json:"password" validate:"required,min=6"`
	Role      string               `bson:"role" json:"role"`
	Phone     string               `bson:"phone" json:"phone" validate:"omitempty,min=12,max=13"`
	Address   string               `bson:"address" json:"address" validate:"omitempty,min=10"`
	Orders    []primitive.ObjectID `bson:"orders" json:"orders"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
}
