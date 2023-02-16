package configs

import (
    "context"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client  {
    Mongo_URI := "mongodb://localhost:27017"
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

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("mysubly").Collection(collectionName)
    return collection
}