package recievers

import (
	get "github.com/jain-chetan/catalog-service/handlers/gethandlers"
)

//RecieverHandler - All Reciever type Handler struct
type RecieverHandler struct {
	GetHandlers *get.GetHandler
}

//Initialization function for RecieverHandler
func Initialization() *RecieverHandler {
	RecieverHandler := new(RecieverHandler)
	RecieverHandler.GetHandlers = new(get.GetHandler)

	return RecieverHandler
}
