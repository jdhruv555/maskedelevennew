package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title       string             `bson:"title" json:"title" validate:"required,min=3"`
	Description string             `bson:"description" json:"description" validate:"required,min=10"`
	Images      []string           `bson:"images" json:"images" validate:"required,min=1"`
	Price       float64            `bson:"price" json:"price" validate:"required,gt=0"`
	Category    string             `bson:"category" json:"category" validate:"required"`
	Sizes       []string           `bson:"sizes" json:"sizes,omitempty"`
	InStock     int                `bson:"inStock" json:"inStock" validate:"required,gte=0"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
