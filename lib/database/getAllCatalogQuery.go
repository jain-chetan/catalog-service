package database

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/jain-chetan/catalog-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllProductsQuery query gets all products based on filters or retrieves all products without filter
func (dc *DBRepo) GetAllProductsQuery(queryParams map[string][]string) ([]model.Catalog, error) {

	var catalogs []model.Catalog
	collection := dc.DBClient.Database("development").Collection("Catalogs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "price", Value: 1}})

	var filter primitive.D
	// cursor, err := collection.Find(ctx, bson.D{primitive.E{Key: "IsDeleted", Value: false}}, findOptions)

	if paramProductName, ok := queryParams["productName"]; ok {
		productName := strings.Join(paramProductName, " ")
		filter = bson.D{primitive.E{Key: "productName",
			Value: primitive.Regex{Pattern: productName, Options: "i"}}, primitive.E{Key: "IsDeleted", Value: false}}
		log.Println("Filter ", filter)
		// cursor, err = collection.Find(ctx, filter, findOptions)
		// if err != nil {
		// 	log.Println("Error finding the data ", err)
		// 	return catalogs, err
		// }
	}

	if paramManufacturer, ok := queryParams["manufacturer"]; ok {
		manufacturer := strings.Join(paramManufacturer, " ")
		filter = bson.D{primitive.E{Key: "manufacturer",
			Value: primitive.Regex{Pattern: manufacturer, Options: "i"}}, primitive.E{Key: "IsDeleted", Value: false}}

		log.Println("Filter ", filter)
		// cursor, err = collection.Find(ctx, filter, findOptions)
		// if err != nil {
		// 	log.Println("Error finding the data ", err)
		// 	return catalogs, err
		// }
	}

	if paramCategory, ok := queryParams["category"]; ok {
		categoryName := strings.Join(paramCategory, " ")
		filter = bson.D{primitive.E{Key: "categoryName",
			Value: primitive.Regex{Pattern: categoryName, Options: "i"}}, primitive.E{Key: "IsDeleted", Value: false}}
		log.Println("Filter ", filter)

		// cursor, err = collection.Find(ctx, filter, findOptions)
		// if err != nil {
		// 	log.Println("Error finding the data ", err)
		// 	return catalogs, err
		// }
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Println("Error finding the data ", err)
		return catalogs, err
	}

	for cursor.Next(ctx) {
		var catalog model.Catalog
		err := cursor.Decode(&catalog)
		if err != nil {
			log.Println("Error in decoding cursor ", err)
			return catalogs, err
		}
		catalogs = append(catalogs, catalog)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return catalogs, err
	}

	return catalogs, nil
}
