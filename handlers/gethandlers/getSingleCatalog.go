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

	pathParam := mux.Vars(request)
	productID := pathParam["productID"]

	log.Println("Path parameter ", productID)

	catalog, err := interfaces.DBClient.GetSingleProductQuery(productID)
	if err != nil {
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
