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
var MongoClient *mongo.Client

// InitMongo initializes MongoDB with optimized connection pooling
func InitMongo() error {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	if dbName == "" {
		dbName = "masked11"
	}

	// Configure connection pool settings
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPI).
		SetMaxPoolSize(100).           // Maximum number of connections in the pool
		SetMinPoolSize(5).             // Minimum number of connections in the pool
		SetMaxConnIdleTime(30 * time.Minute). // Maximum time a connection can remain idle
		SetRetryWrites(true).          // Enable retryable writes
		SetRetryReads(true).           // Enable retryable reads
		SetServerSelectionTimeout(5 * time.Second). // Server selection timeout
		SetSocketTimeout(30 * time.Second).         // Socket timeout
		SetConnectTimeout(10 * time.Second)        // Connection timeout

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Test the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	MongoClient = client
	Mongo = client.Database(dbName)

	// Create indexes for optimal performance
	if err := ensureIndexes(); err != nil {
		log.Printf("Warning: Index creation failed: %v", err)
	}

	// Start connection monitoring
	go monitorConnection()

	log.Println("✅ MongoDB connected with optimized settings")
	return nil
}

// ensureIndexes creates necessary indexes for performance
func ensureIndexes() error {
	// User indexes
	userColl := Mongo.Collection("users")
	userIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("email_unique"),
		},
		{
			Keys: bson.D{{Key: "role", Value: 1}},
			Options: options.Index().SetName("role_index"),
		},
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
			Options: options.Index().SetName("created_at_index"),
		},
	}

	// Product indexes
	productColl := Mongo.Collection("products")
	productIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "category", Value: 1}},
			Options: options.Index().SetName("category_index"),
		},
		{
			Keys: bson.D{{Key: "price", Value: 1}},
			Options: options.Index().SetName("price_index"),
		},
		{
			Keys: bson.D{{Key: "createdAt", Value: -1}},
			Options: options.Index().SetName("created_at_index"),
		},
		{
			Keys: bson.D{
				{Key: "title", Value: "text"},
				{Key: "description", Value: "text"},
			},
			Options: options.Index().SetName("text_search_index"),
		},
		{
			Keys: bson.D{
				{Key: "category", Value: 1},
				{Key: "price", Value: 1},
			},
			Options: options.Index().SetName("category_price_index"),
		},
		{
			Keys: bson.D{
				{Key: "inStock", Value: 1},
				{Key: "category", Value: 1},
			},
			Options: options.Index().SetName("stock_category_index"),
		},
	}

	// Create user indexes
	for _, index := range userIndexes {
		_, err := userColl.Indexes().CreateOne(context.TODO(), index)
		if err != nil {
			log.Printf("Warning: Failed to create user index: %v", err)
		}
	}

	// Create product indexes
	for _, index := range productIndexes {
		_, err := productColl.Indexes().CreateOne(context.TODO(), index)
		if err != nil {
			log.Printf("Warning: Failed to create product index: %v", err)
		}
	}

	return nil
}

// monitorConnection monitors MongoDB connection health
func monitorConnection() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		
		if err := MongoClient.Ping(ctx, readpref.Primary()); err != nil {
			log.Printf("⚠️ MongoDB connection health check failed: %v", err)
		}
		
		cancel()
	}
}

// GetMongoStats returns MongoDB connection statistics
func GetMongoStats() map[string]interface{} {
	if MongoClient == nil {
		return map[string]interface{}{
			"status": "disconnected",
		}
	}

	// Get connection pool statistics
	stats := MongoClient.NumberSessionsInProgress()
	
	return map[string]interface{}{
		"status":           "connected",
		"sessionsInProgress": stats,
		"database":         Mongo.Name(),
		"timestamp":        time.Now().UTC(),
	}
}

// CloseMongo closes the MongoDB connection
func CloseMongo() error {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return MongoClient.Disconnect(ctx)
	}
	return nil
}

// GetCollection returns a MongoDB collection with optimized settings
func GetCollection(name string) *mongo.Collection {
	return Mongo.Collection(name)
}

// ExecuteWithTimeout executes a MongoDB operation with timeout
func ExecuteWithTimeout(ctx context.Context, timeout time.Duration, operation func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return operation(ctx)
}

// BulkWrite performs bulk write operations with optimization
func BulkWrite(collection *mongo.Collection, operations []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := options.BulkWrite().SetOrdered(false) // Unordered for better performance
	return collection.BulkWrite(ctx, operations, opts)
}

// AggregateWithOptions performs aggregation with optimized options
func AggregateWithOptions(collection *mongo.Collection, pipeline interface{}, opts *options.AggregateOptions) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if opts == nil {
		opts = options.Aggregate()
	}
	
	// Set default optimization options
	opts.SetAllowDiskUse(true)
	opts.SetBatchSize(1000)

	return collection.Aggregate(ctx, pipeline, opts)
}

// FindWithOptions performs find operations with optimized options
func FindWithOptions(collection *mongo.Collection, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if opts == nil {
		opts = options.Find()
	}

	// Set default optimization options
	opts.SetBatchSize(1000)
	opts.SetNoCursorTimeout(false)

	return collection.Find(ctx, filter, opts)
}

// FindOneWithOptions performs findOne operations with optimized options
func FindOneWithOptions(collection *mongo.Collection, filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.FindOne(ctx, filter, opts)
}

// InsertOneWithOptions performs insertOne operations with optimized options
func InsertOneWithOptions(collection *mongo.Collection, document interface{}, opts *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.InsertOne(ctx, document, opts)
}

// UpdateOneWithOptions performs updateOne operations with optimized options
func UpdateOneWithOptions(collection *mongo.Collection, filter interface{}, update interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.UpdateOne(ctx, filter, update, opts)
}

// DeleteOneWithOptions performs deleteOne operations with optimized options
func DeleteOneWithOptions(collection *mongo.Collection, filter interface{}, opts *options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.DeleteOne(ctx, filter, opts)
}
