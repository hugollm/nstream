package main

import (
	"log"
	"net/http"
	"nstream/api"
	"nstream/login"
	"nstream/logout"
	notesnew "nstream/notes/new"
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

var nts api.Api = api.NewApi([]api.Endpoint{
	status.Status{},
	signup.Signup{},
	login.Login{},
	logout.Logout{},
	notesnew.NewNote{},
})
