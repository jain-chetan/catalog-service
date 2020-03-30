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

	var catalog model.Catalog
	pathUpdateParam := mux.Vars(request)
	productID := pathUpdateParam["productID"]

	log.Println("Path param for update ", productID)

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

	//updateResponse, err := interfaces.DBClient.UpdateProductQuery(productID, catalog)
	var err error
	if err != nil {
		log.Println("Error in Updating Data in database ", err)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Decoding JSON Body",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	//TODO: Condition goes here if updatedField count value is 0 after DB Call

	result := model.Response{
		Code:    200,
		Message: "Ok",
	}
	response.Header().Add("Status", "200")
	json.NewEncoder(response).Encode(result)
}
