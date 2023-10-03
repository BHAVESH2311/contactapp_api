package usercontroller

import (
	user_service "contactapp_api/components/User/user_service"
	auth "contactapp_api/middleware"
	"contactapp_api/validators"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type Info struct {
	Username    string
	Password string
}

func handlePanic(err interface{}){
	log.Printf("error %V",err)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user *user_service.User

	json.NewDecoder(r.Body).Decode(&user)

	if !validators.ValidateName(user.Email) {
		panic("Please enter a valid username")
	}

	newUser := user_service.CreateUser(user.Name, user.Email, user.Password)

	var claims = &auth.Claims{
		Name:     newUser.Name,
		IsAdmin:  false,
		IsActive: false,
	}

	token, err := auth.SignJWT(*claims)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "authentication",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 5),
		Secure:  true,
	})

	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Created Successfully"))
}

func Signin(w http.ResponseWriter, r *http.Request) {
	defer func() {
        if r := recover(); r != nil {
            handlePanic(r)
        }
    }()
	w.Header().Set("ContentType", "application/json")
	var enteredInfo Info
	err := json.NewDecoder(r.Body).Decode(&enteredInfo)
	if err != nil {
		panic(err)
	}

	user, err := user_service.GetUserByEmail(enteredInfo.Username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid Username"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(enteredInfo.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials"))
		return
	}

	var claims = &auth.Claims{
		Name:     user.Name,
		IsAdmin:  false,
		IsActive: false,
	}

	token, err := auth.SignJWT(*claims)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "authentication",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 5),
		Secure:  true,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Logged In Successfully"))
}

