package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/catalog-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetSingleProductQuery gets product details based on supplied productID
func (dc *DBRepo) GetSingleProductQuery(productID string) (model.Catalog, error) {
	collection := dc.DBClient.Database("development").Collection("Catalogs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//converting string ID into primitive hexadecimal object
	ID, err := primitive.ObjectIDFromHex(productID)
	var catalog model.Catalog

	//Call to the database to fetch the data
	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: ID},
		primitive.E{Key: "IsDeleted", Value: false}}).Decode(&catalog)

	// error handling
	if err != nil {
		log.Println("Error in get Single Product ", err)
		return catalog, err
	}

	//Returning the data
	return catalog, nil

}
