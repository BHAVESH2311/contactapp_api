package contactdetails

import (
	"github.com/gorilla/mux"
	contact_details_controller "contactapp_api/components/contact_details/contact_details_controller"
)

func ContactDetailRouter(router *mux.Router) *mux.Router {
	contactDetailsdRouter := router.PathPrefix("/user/{userId}/contact/{contactId}/contactdetail").Subrouter()
	contactDetailsdRouter.HandleFunc("", contact_details_controller.CreateContactDetail).Methods("POST")
	contactDetailsdRouter.HandleFunc("/{contactDetailId}", contact_details_controller.UpdateContactDetail).Methods("PUT")
	contactDetailsdRouter.HandleFunc("/{contactDetailId}", contact_details_controller.DeleteContactDetail).Methods("DELETE")

	return contactDetailsdRouter
}