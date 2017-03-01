package main

import (
	"log"
	"net/http"
	"simple-json-restapi/routers"
)

func main() {
	//init object
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
