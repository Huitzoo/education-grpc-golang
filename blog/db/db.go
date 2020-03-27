package db

import (
	"blog/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var url string = "mongodb+srv://<username>:<password>@<cluster>-5bevl.mongodb.net/test?retryWrites=true&w=majority"

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
func (db *DB) InsertOne(collection int, data interface{}) (string, error) {
	result, err := db.collections[collection].InsertOne(context.Background(), data)
	if err != nil {
		return "", err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", fmt.Errorf("Can't cast oid to objectid")
	}

	return oid.String(), nil
}

//UpdateOne is
func (db *DB) UpdateOne(collection int, iod string, data interface{}) (*mongo.UpdateResult, error) {
	blog, errFind := db.FindByID(0, iod)

	if errFind != nil {
		return nil, errFind
	}
	blogModel := &models.ItemBlog{}
	blog.Decode(blogModel)

	parseData := data.(models.ItemBlog)

	updateData := parseData.CreateUpdateBlog(blogModel)

	iodHex, errCast := primitive.ObjectIDFromHex(iod)
	if errCast != nil {
		return nil, errCast
	}

	filter := bson.M{"_id": bson.M{"$eq": iodHex}}
	update := bson.M{"$set": updateData}

	result, errUpdate := db.collections[collection].UpdateOne(context.Background(), filter, update)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return result, nil
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
