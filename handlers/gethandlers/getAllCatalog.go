package gethandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jain-chetan/catalog-service/interfaces"
	"github.com/jain-chetan/catalog-service/model"
)

//GetAllProductsHandler handles request and response to get all products
func (getAllData *GetHandler) GetAllProductsHandler(response http.ResponseWriter, request *http.Request) {

	queryParams := request.URL.Query()
	var catalogs []model.Catalog

	log.Println(queryParams)
	response.Header().Add("content-type", "application/json")

	catalogs, err := interfaces.DBClient.GetAllProductsQuery(queryParams)
	log.Println("Catalogs recieved ", catalogs)
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

	if len(catalogs) <= 0 {
		errResponse := model.Response{
			Code:    400,
			Message: "No Records Found",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	response.Header().Add("Status", "200")
	json.NewEncoder(response).Encode(catalogs)

}
