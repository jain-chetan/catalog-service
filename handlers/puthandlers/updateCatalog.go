package puthandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/catalog-service/interfaces"
	"github.com/jain-chetan/catalog-service/model"
)

type PutHandler struct{}

func (updateData *PutHandler) UpdateProductHandler(response http.ResponseWriter, request *http.Request) {

	// var catalog model.Catalog
	pathUpdateParam := mux.Vars(request)
	productID := pathUpdateParam["productID"]

	log.Println("Path param for update ", productID)

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Call to Database to get product data
	catalog, err := interfaces.DBClient.GetSingleProductQuery(productID)
	if err != nil {
		log.Println("Error in getting data ", err)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in getting data",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	//Decoding to same catalog model
	errDecode := json.NewDecoder(request.Body).Decode(&catalog)

	if errDecode != nil {
		log.Println("Error in Decoding JSON Body")
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Decoding JSON Body",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	isProductExist := interfaces.DBClient.CheckProductExist(productID)
	if !isProductExist {
		errResponse := model.Response{
			Code:    400,
			Message: "No Records Found",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	modifiedCount, err := interfaces.DBClient.UpdateProductQuery(productID, catalog)

	if err != nil || modifiedCount <= 0 {
		log.Println("Error in Updating Data in database ", err)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Updating Data",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	result := model.Response{
		Code:    200,
		Message: "Ok",
	}
	response.Header().Add("Status", "200")
	json.NewEncoder(response).Encode(result)
}
