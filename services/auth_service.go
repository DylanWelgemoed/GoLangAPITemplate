package services

import (
	"encoding/json"
	"net/http"
	"github.com/DylanWelgemoed/GoLangAPITemplate/api/parameters"
	"github.com/DylanWelgemoed/GoLangAPITemplate/core/authentication"
	"github.com/DylanWelgemoed/GoLangAPITemplate/models"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

func Login(requestUser *models.UserLogin) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.UUID)
		  
  		if err != nil {
   			return http.StatusInternalServerError, []byte("")
  		} else {
   			response, _ := json.Marshal(parameters.TokenAuthentication{token})
   			return http.StatusOK, response
  		}
	}
	 
	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.UserLogin) []byte {
 	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.UUID)
	 
 	if err != nil {
 		panic(err)
	}
	
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	 
 	if err != nil {
		panic(err)
	}
	 
 	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	 
	if err != nil {
		return err
	}
	
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}