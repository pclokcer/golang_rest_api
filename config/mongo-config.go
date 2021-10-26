package config

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Database {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Env file error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION")))

	if err != nil {
		panic(err)
	}

	println("Mongo Bağlandı")

	mongoDB := client.Database("helorobo_db")

	return mongoDB
}
