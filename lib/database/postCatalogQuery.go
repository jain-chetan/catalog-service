package database

import (
	"context"
	"time"

	"github.com/jain-chetan/catalog-service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateProductsQuery query function to insert products in database
func (dc *DBRepo) CreateProductsQuery(catalog model.Catalog) (model.CreateResponse, error) {
	var result model.CreateResponse
	emptyResponse := model.CreateResponse{}
	collection := dc.DBClient.Database("development").Collection("Catalogs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, catalog)
	if err != nil {
		return emptyResponse, err
	}

	result.ID = res.InsertedID.(primitive.ObjectID)

	return result, nil
}
