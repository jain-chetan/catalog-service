package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CheckProductExist checks whether product exist or deleted
func (dc *DBRepo) CheckProductExist(productID string) bool {
	collection := dc.DBClient.Database("development").Collection("Catalogs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var count int64
	ID, err := primitive.ObjectIDFromHex(productID)

	//Getting the count of documents for the supplied ID
	count, err = collection.CountDocuments(ctx, bson.D{primitive.E{Key: "_id", Value: ID},
		primitive.E{Key: "IsDeleted", Value: false}})

	log.Println("Total count matched ", count)
	if count <= 0 || err != nil {
		if err != nil {
			log.Println("Check product exist error ", err)
		}
		return false
	}

	return true
}
