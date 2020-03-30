package posthandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jain-chetan/catalog-service/interfaces"
	"github.com/jain-chetan/catalog-service/model"
)

//PostHandler structure to group all get methods
type PostHandler struct{}

//PostCatalogHandler to handle JSON request and response for database insertion of products
func (postData *PostHandler) PostCatalogHandler(response http.ResponseWriter, request *http.Request) {
	//Catalog model to decode JSON data
	var catalog model.Catalog

	//Decoding JSON Body from Request
	errDecode := json.NewDecoder(request.Body).Decode(&catalog)

	//Error Handling for Decoding JSON Request
	if errDecode != nil {
		log.Println("Error in Decoding JSON Body", errDecode)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Decoding JSON Body",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	//Calling CreateProductsQuery to insert products in database
	result, err := interfaces.DBClient.CreateProductsQuery(catalog)
	if err != nil {
		log.Println("Error in inserting data ", err)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Inserting Data",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	//Returning created ID back for new product inserted.
	log.Println("Inserted ID is ", result.ID)
	result = model.CreateResponse{
		ID:      result.ID,
		Code:    201,
		Message: "Successfully Created",
	}
	response.Header().Add("Status", "201")
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(response).Encode(result)
}
