package main

import (
	"log"
	"net/http"
	"nstream/api"
	"nstream/signup"
	"nstream/status"
)

func main() {
	http.HandleFunc("/", handler)
	log.SetFlags(log.Flags() + log.LUTC)
	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(response http.ResponseWriter, request *http.Request) {
	nts.Handle(request, response)
}

var nts api.NtsApi = api.NewApi([]api.Endpoint{
	status.Status{},
	signup.Signup{},
})
