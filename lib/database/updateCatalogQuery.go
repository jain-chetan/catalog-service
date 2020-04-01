package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/catalog-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dc *DBRepo) UpdateProductQuery(productID string, catalog model.Catalog) (int64, error) {

	collection := dc.DBClient.Database("development").Collection("Catalogs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//converting string ID into primitive hexadecimal object
	ID, errConversion := primitive.ObjectIDFromHex(productID)
	if errConversion != nil {
		return 0, errConversion
	}

	res, errUpdate := collection.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: ID}},
		bson.D{primitive.E{Key: "$set", Value: catalog}})

	log.Println("Update filter and data ", bson.D{primitive.E{Key: "_id", Value: ID}},
		bson.D{primitive.E{Key: "$set", Value: catalog}})

	if errUpdate != nil {
		return 0, errUpdate
	}
	log.Println("modifiedCount ", res.ModifiedCount)
	log.Println("matchedCount ", res.MatchedCount)
	log.Println("Total ", res.ModifiedCount+res.MatchedCount)
	return (res.ModifiedCount + res.MatchedCount), nil

}
