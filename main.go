package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Anniversary struct {
	Batchcomplete string `json:"batchcomplete"`
	Warnings      struct {
		Extracts struct {
			message string `json:"*"`
		} `json:"extracts"`
	} `json:"warnings"`
	Query struct {
		Pages struct {
			Id struct {
				Pageid  int    `json:"pageid"`
				Ns      int    `json:"ns"`
				Title   string `json:"title"`
				Extract string `json:"extract"`
			} `json:"456328"`
		} `json:"pages"`
	} `json:"query"`
}

const (
	ENDPOINT = "https://ja.wikipedia.org/w/api.php?action=query&prop=extracts&titles=%e6%97%a5%e6%9c%ac%e3%81%ae%e8%a8%98%e5%bf%b5%e6%97%a5%e4%b8%80%e8%a6%a7&format=xml"
)

func Search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", ENDPOINT, nil)

	response, _ := client.Do(request) // Endpointにgetを投げる
	defer response.Body.Close()       //deferで終了時に実行

	body, _ := ioutil.ReadAll(response.Body)

	var anniversaries []Anniversary
	json.Unmarshal(body, &anniversaries)

	//fmt.Fprintln(w, string(body))

	responseString := string(body)
	punctuateMonthly := regexp.MustCompile(`(\d{1,2})日\s-\s(.+日|.+デー)`)

	punctuatedByMonth := punctuateMonthly.FindAllString(responseString, 400)

	for s1, s2 := range punctuatedByMonth {
		fmt.Fprintln(w, s1+1, s2)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/search/:month/:date", Search)
	log.Fatal(http.ListenAndServe(":8080", router))
}
