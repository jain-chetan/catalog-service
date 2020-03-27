package recievers

import (
	delete "github.com/jain-chetan/catalog-service/handlers/deletehandlers"
	get "github.com/jain-chetan/catalog-service/handlers/gethandlers"
	post "github.com/jain-chetan/catalog-service/handlers/posthandlers"
	put "github.com/jain-chetan/catalog-service/handlers/puthandlers"
)

//RecieverHandler - All Reciever type Handler struct
type RecieverHandler struct {
	GetHandlers    *get.GetHandler
	PostHandlers   *post.PostHandler
	PutHandlers    *put.PutHandler
	DeleteHandlers *delete.DeleteHandler
}

//Initialization function for RecieverHandler
func Initialization() *RecieverHandler {
	RecieverHandler := new(RecieverHandler)
	RecieverHandler.GetHandlers = new(get.GetHandler)
	RecieverHandler.PostHandlers = new(post.PostHandler)
	RecieverHandler.PutHandlers = new(put.PutHandler)
	RecieverHandler.DeleteHandlers = new(delete.DeleteHandler)

	return RecieverHandler
}
