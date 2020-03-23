package gethandlers

import (
	"encoding/json"
	"net/http"

	"github.com/jain-chetan/catalog-service/model"
)

//GetHandler structure to group all get methods
type GetHandler struct{}

//PingHandler to check catalog service response
func (getData *GetHandler) PingHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	pingResponse := model.Response{
		Code:    200,
		Message: "Ok",
	}

	json.NewEncoder(response).Encode(pingResponse)

}
