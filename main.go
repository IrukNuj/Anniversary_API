package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)



const (
	ENDPOINT = "https://ja.wikipedia.org/w/api.php?action=query&prop=extracts&titles=%e6%97%a5%e6%9c%ac%e3%81%ae%e8%a8%98%e5%bf%b5%e6%97%a5%e4%b8%80%e8%a6%a7&format=json"
)

func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, _ := http.Get(ENDPOINT) // Endpointにgetを投げる
	defer response.Body.Close() //deferで終了時に実行

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Fprintln(w, body)

}

func main() {
	router := httprouter.New()
	router.GET("/", Search)
	log.Fatal(http.ListenAndServe(":8080", router))
}
