package controllers

import (
    "encoding/json"
	"net/http"
	"strconv"
	"io"
	"io/ioutil"
    "github.com/gorilla/mux"
	"github.com/DylanWelgemoed/GoLangAPITemplate/structs"
	"github.com/DylanWelgemoed/GoLangAPITemplate/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
    if err := json.NewEncoder(w).Encode(services.GetUsers()); err != nil {
        panic(err)
    }
}

func FindUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	var userId int
	var err error

	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		panic(err)
	}

	user := services.FindUser(userId)

	if user.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	if err := json.NewEncoder(w).Encode(structs.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New User", "email":"newuser@test.com"}' http://localhost:5000/users
*/
func CreateUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var user structs.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	u := services.CreateUser(user)
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}