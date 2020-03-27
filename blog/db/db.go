package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var url string = "mongodb+srv://<user>:<password>@<cluster>-5bevl.mongodb.net/test?retryWrites=true&w=majority"

//DB is
type DB struct {
	client      *mongo.Client
	collections []*mongo.Collection
}

//ConnectDB is
func (db *DB) ConnectDB() {
	var err error

	db.client, err = mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		log.Fatal(err)
	}
	db.client.Connect(context.TODO())
	db.collections = append(db.collections, db.client.Database("education").Collection("blog"))
	fmt.Println("connected mongodb")
}

//CloseConnnection is
func (db *DB) CloseConnnection() {
	db.client.Disconnect(context.TODO())
}

//InsertOne is
func (db *DB) InsertOne(collection int, data interface{}) (*mongo.InsertOneResult, error) {
	return db.collections[collection].InsertOne(context.Background(), data)
}

//FindByID is
func (db *DB) FindByID(collection int, id string) (*mongo.SingleResult, error) {
	iod, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": iod}

	return db.collections[collection].FindOne(context.Background(), filter), nil
}
