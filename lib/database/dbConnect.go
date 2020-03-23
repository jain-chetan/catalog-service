package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/catalog-service/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBRepo satisfies the interface by implementing all the methods
type DBRepo struct {
	DBClient *mongo.Client
}

//DBConnect Method to connect to Db
func (dc *DBRepo) DBConnect(config model.DBConfig) error {
	var err error
	// Format DB configs to required format connect DB
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//mongodb://localhost:27017
	dc.DBClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.User+"://"+config.Host+":"+config.Port))
	if err != nil {
		log.Printf("Unable to connect DB %v", err)
		return err
	}

	log.Printf("MongoDB started at %s PORT", config.Port)
	return err
}
