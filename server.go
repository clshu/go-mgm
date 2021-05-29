package main

import (
	"net/http"

	"github.com/clshu/go-mgm/api"
	"github.com/clshu/go-mgm/dbmgm"
	"github.com/gorilla/mux"
)

func main() {

	err := dbmgm.Connect()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/user/create", api.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", api.LogIn).Methods("POST")
	http.ListenAndServe(":3030", router)

}
