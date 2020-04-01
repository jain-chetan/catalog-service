package deletehandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/catalog-service/interfaces"
	"github.com/jain-chetan/catalog-service/model"
)

type DeleteHandler struct{}

func (deleteData *DeleteHandler) DeleteProductHandler(response http.ResponseWriter, request *http.Request) {

	deletePathParam := mux.Vars(request)
	productID := deletePathParam["productID"]

	log.Println("Path Param ", productID)

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

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

	deleteResponse, err := interfaces.DBClient.DeleteProductQuery(productID)
	if err != nil || deleteResponse <= 0 {
		log.Println("Error in deleting Data in database ", err)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Deleting data",
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
