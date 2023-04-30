package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var SubsCollection *mongo.Collection

func init() {
	// Client instance
	DB = ConnectDB()
	SubsCollection = GetCollection("subs")
}

func ConnectDB() *mongo.Client {
	Mongo_URI := "mongodb+srv://saintmalik:wCxDL4bJnBscJii9@mysublycluster.qzh4zzx.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	collection := DB.Database("mysubly").Collection(collectionName)
	return collection
}
