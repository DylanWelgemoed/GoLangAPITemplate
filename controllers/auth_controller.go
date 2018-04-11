package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/DylanWelgemoed/GoLangAPITemplate/services"
	"github.com/DylanWelgemoed/GoLangAPITemplate/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.UserLogin)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)

 	w.Header().Set("Content-Type", "application/json")
 	w.WriteHeader(responseStatus)
 	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(models.UserLogin)
	 
 	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	 
	w.Header().Set("Content-Type", "application/json")
 	w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	 
	w.Header().Set("Content-Type", "application/json")
	 
 	if err != nil {
  		w.WriteHeader(http.StatusInternalServerError)
 	} else {
  		w.WriteHeader(http.StatusOK)
 	}
}