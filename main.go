package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Search() {

}

func main() {
	router := httprouter.New()
	router.GET("/", Search)
	log.Fatal(http.ListenAndServe(":8080", router))

}
