package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/clshu/go-mgm/api"
	"github.com/clshu/go-mgm/dbmgm"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	setUpEnv()

	err := dbmgm.Connect()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/user/create", api.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", api.LogIn).Methods("POST")
	http.ListenAndServe(":3030", router)

}

// Set up development or test environment variables
// Expect files to be in project_root/config
// Set GOGO_ENV first before go run server.go
func setUpEnv() {
	const dir string = "config/"
	var fname string
	gogo := os.Getenv("GOGO_ENV")

	switch gogo {
	case "dev":
		fname = dir + "dev.env"
		break
	case "test":
		fname = dir + "test.env"
		break
	default:
		// production environment
		// Do nothing. Let clound platform environment take over
		return
	}

	envMap, err := godotenv.Read(fname)
	if err != nil {
		fmt.Printf("Reading file %v failed. %v", fname, err.Error())
		return
	}

	for key, value := range envMap {
		os.Setenv(key, value)
		// fmt.Printf("%v=%v\n", key, os.Getenv(key))
	}

}
