package main

import (
	"log"
	"net/http"
	"simple-json-restapi/routers"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
