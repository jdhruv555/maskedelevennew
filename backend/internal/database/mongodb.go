package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mongo *mongo.Database

func InitMongo() {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	Mongo = client.Database(dbName)
	log.Println("MongoDB connected")

	if err := ensureUserIndexes(); err != nil {
		log.Fatalf("Index creation failed: %v", err)
	}
}

func ensureUserIndexes() error {
	userColl := Mongo.Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := userColl.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
