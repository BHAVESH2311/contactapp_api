package usercontroller

import (
	user_service "contactapp_api/components/User/user_service"
	"contactapp_api/validators"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user_service.User
	json.NewDecoder(r.Body).Decode(&user)

	if !validators.ValidateName(user.Email) {
		panic("Please enter a valid username")
	}

	var newUser = user_service.CreateAdmin(user.Name,user.Email,user.Password)
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(newUser)

}

func ReadAllUsers(w http.ResponseWriter, r *http.Request) {

	users := user_service.GetAllUsers()

	w.Header().Set("ContentType", "application/json")

	json.NewEncoder(w).Encode(map[string][]user_service.User{"All-Users": users})
}


func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)

	user, err := user_service.GetUserById(mp["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	var container *user_service.User
	json.NewDecoder(r.Body).Decode(&container)

	user_service.UpdateUserById(container, user)

	w.Header().Set("ContentType", "application/json")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"Action":"User Updated Successfully"})

}


func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)

	_, err := user_service.GetUserById(mp["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	user_service.DeleteUserByID(mp["id"])
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Action": "User Successfully Deleted"})
}


func GetUserById(w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)

	users, err := user_service.GetUserById(mp["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	w.Header().Set("ContentType", "application/json")

	json.NewEncoder(w).Encode(users)
}

