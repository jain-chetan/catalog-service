package gethandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/catalog-service/interfaces"
	"github.com/jain-chetan/catalog-service/model"
)

//GetSingleProductHandler to handle request and response for get single product
func (getData *GetHandler) GetSingleProductHandler(response http.ResponseWriter, request *http.Request) {

	//Getting productID from path parameter
	pathParam := mux.Vars(request)
	productID := pathParam["productID"]

	log.Println("Path parameter ", productID)

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Checking if product Exist before getting products for that ID
	isProductExist := interfaces.DBClient.CheckProductExist(productID)
	log.Println("IsProductExist", isProductExist)
	if !isProductExist {
		errResponse := model.Response{
			Code:    400,
			Message: "No Records Found",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

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

	response.Header().Add("Status", "200")
	json.NewEncoder(response).Encode(catalog)
}
