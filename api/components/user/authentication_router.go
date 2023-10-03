package user

import (
	userauthentication "contactapp_api/components/user/user_controller"

	"github.com/gorilla/mux"
)

func AuthenticationRouter(router *mux.Router) *mux.Router {
	authenticationRouter := router.PathPrefix("/authentication").Subrouter()
	authenticationRouter.HandleFunc("/register", userauthentication.Register).Methods("POST")
	authenticationRouter.HandleFunc("/login", userauthentication.Signin).Methods("POST")
	return authenticationRouter
}