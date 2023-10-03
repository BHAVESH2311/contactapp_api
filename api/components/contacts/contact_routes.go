package contact

import (
	contact_controller "contactapp_api/components/contacts/contact_controller"

	"github.com/gorilla/mux"
)

func ContactRouter(router *mux.Router) *mux.Router {
	contactRouter := router.PathPrefix("/user/{userId}/contact").Subrouter()

	contactRouter.HandleFunc("", contact_controller.GetContactsForUser).Methods("GET")
	contactRouter.HandleFunc("", contact_controller.AddNewContact).Methods("POST")
	contactRouter.HandleFunc("/{contactId}", contact_controller.UpdateExistingContact).Methods("PUT")
	contactRouter.HandleFunc("/{contactId}", contact_controller.DeleteContact).Methods("DELETE")

	return contactRouter
}