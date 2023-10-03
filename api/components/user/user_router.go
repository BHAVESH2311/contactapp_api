package user

import (
	 user_controller "contactapp_api/components/user/user_controller"
	 "contactapp_api/middleware"
	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router) *mux.Router {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(authentication.AuthenticationHandler)
	userRouter.HandleFunc("", user_controller.ReadAllUsers).Methods("GET")
	userRouter.HandleFunc("", user_controller.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id}", user_controller.DeleteUserByID).Methods("DELETE")
	userRouter.HandleFunc("/{id}", user_controller.UpdateUserById).Methods("PUT")
	userRouter.HandleFunc("/{id}", user_controller.GetUserById).Methods("GET")
	return userRouter
}