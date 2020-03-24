package recievers

import (
	get "github.com/jain-chetan/catalog-service/handlers/gethandlers"
	post "github.com/jain-chetan/catalog-service/handlers/posthandlers"
)

//RecieverHandler - All Reciever type Handler struct
type RecieverHandler struct {
	GetHandlers  *get.GetHandler
	PostHandlers *post.PostHandler
}

//Initialization function for RecieverHandler
func Initialization() *RecieverHandler {
	RecieverHandler := new(RecieverHandler)
	RecieverHandler.GetHandlers = new(get.GetHandler)
	RecieverHandler.PostHandlers = new(post.PostHandler)

	return RecieverHandler
}
