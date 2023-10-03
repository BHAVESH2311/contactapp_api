package authentication

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Id      string
	Name     string
	IsAdmin  bool
	IsActive bool
	jwt.StandardClaims
}

var secret = []byte("JWT Token")

func SignJWT(claims Claims) (string, error) {


	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * 5).Unix()


	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


	signToken, err := tokenObj.SignedString(secret)
	if err != nil {
		return "Token Invalid", err
	}
	return signToken, nil 
}

func TokenVerify(token string) (*Claims, error) {
	var userClaim = &Claims{}
	tokenObj, err := jwt.ParseWithClaims(token, userClaim, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Access Denied")
		}
		return nil, errors.New("status : bad request")
	}
	if !tokenObj.Valid {
		return nil, errors.New(" Invalid Token ")
	}
	return userClaim, nil
}

func AuthenticationHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("ContentType", "application/json")
		token := r.Header.Values("authentication")
		if len(token) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(" Token does not exist"))
			return
		}
		tokenStr := token[0]
		_, err := TokenVerify(tokenStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Access Denied"))
			return
		}
		h.ServeHTTP(w, r)
	})
}