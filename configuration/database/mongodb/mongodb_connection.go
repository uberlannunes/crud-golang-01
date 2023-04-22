package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	mongodb_uri := os.Getenv(MONGODB_URL)
	monodb_users_db := os.Getenv(MONGODB_USER_DB)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodb_uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(monodb_users_db), nil
}

func InitConnection() {

	mongodb_uri := os.Getenv(MONGODB_URL)
	// monodb_users_db := os.Getenv(MONGODB_USER_DB)
	// mongodb_user := os.Getenv(MONGODB_USER)
	// mongodb_password := os.Getenv(MONGODB_PASSWORD)

	ctx := context.Background()
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI("mongodb+srv://ubnunes-mongodb:cvvKYxlzSQY5ODg5@cluster0.fmgoovh.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", mongodb_user, mongodb_password, mongodb_uri)).SetServerAPIOptions(serverAPI)
	opts := options.Client().ApplyURI(mongodb_uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	// if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
	// if err := client.Database(monodb_users_db).RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
	// 	panic(err)
	// }

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
