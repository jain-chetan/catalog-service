package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dc *DBRepo) DeleteProductQuery(productID string) (int64, error) {
	collection := dc.DBClient.Database("development").Collection("Catalogs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//converting string ID into primitive hexadecimal object
	ID, errConversion := primitive.ObjectIDFromHex(productID)
	if errConversion != nil {
		return 0, errConversion
	}

	res, errUpdate := collection.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: ID}},
		bson.D{primitive.E{Key: "IsDeleted", Value: true}})

	if errUpdate != nil {
		return 0, errUpdate
	}
	log.Println("modifiedCount ", res.ModifiedCount)

	return res.ModifiedCount, nil
}
