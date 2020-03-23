package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jain-chetan/catalog-service/interfaces"
	"github.com/jain-chetan/catalog-service/lib/database"
	"github.com/jain-chetan/catalog-service/model"
	apiServices "github.com/jain-chetan/catalog-service/recievers"

	"github.com/gorilla/mux"
)

func main() {
	err := initDBClient()
	if err != nil {
		log.Fatal("DB Driver error", err)
	}
	api := apiServices.Initialization()
	router := mux.NewRouter()
	router.HandleFunc("/catalog/ping", api.GetHandlers.PingHandler).Methods("GET")

	http.ListenAndServe(":8082", router)

}

func initDBClient() error {

	var config model.DBConfig

	//Read DB credentials from environment variables
	config.User = os.Getenv("DBUSER")
	config.Port = os.Getenv("PORT")
	config.Host = os.Getenv("HOST")
	interfaces.DBClient = new(database.DBRepo)
	err := interfaces.DBClient.DBConnect(config)
	return err
}
