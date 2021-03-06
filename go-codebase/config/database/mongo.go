package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB return mongo db read & write instance
func InitMongoDB(ctx context.Context) (read *mongo.Database, write *mongo.Database) {
	// init write mongodb
	hostWrite := os.Getenv("MONGO_DB_HOST_WRITE")
	dbNameWrite := os.Getenv("MONGO_DB_NAME_WRITE")
	client, err := mongo.NewClient(options.Client().ApplyURI(hostWrite))
	if err != nil {
		panic(err)
	}
	if err := client.Connect(ctx); err != nil {
		panic(err)
	}
	write = client.Database(dbNameWrite)

	// init read mongodb
	hostRead := os.Getenv("MONGO_DB_HOST_READ")
	dbNameRead := os.Getenv("MONGO_DB_NAME_READ")
	client, err = mongo.NewClient(options.Client().ApplyURI(hostRead))
	if err != nil {
		panic(err)
	}
	if err := client.Connect(ctx); err != nil {
		panic(err)
	}
	read = client.Database(dbNameRead)

	return
}
