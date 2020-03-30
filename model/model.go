package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Catalog structure for catalog document
type Catalog struct {
	ID               primitive.ObjectID `json:"productID,omitempty" bson:"_id,omitempty"`
	ProductName      string             `json:"productName" bson:"productName"`
	ProductImageLink string             `json:"productImageLink" bson:"productImageLink"`
	Manufacturer     string             `json:"manufacturer" bson:"manufacturer"`
	CategoryName     string             `json:"categoryName" bson:"categoryName"`
	Price            float64            `json:"price" bson:"price"`
	IsDeleted        bool               `json:"-" bson:"IsDeleted"`
}

//DBConfig has information required to connect to DB
type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

//Response defines response code and message
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//CreateResponse for insertion response
type CreateResponse struct {
	ID      primitive.ObjectID `json:"productID"`
	Code    int                `json:"code"`
	Message string             `json:"message"`
}
