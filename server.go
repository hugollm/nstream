package main

import (
	"log"
	"net/http"
	"nstream/api"
	"nstream/auth/login"
	"nstream/auth/logout"
	"nstream/auth/signup"
	"nstream/notes/deletenote"
	"nstream/notes/fetchnotes"
	"nstream/notes/newnote"
	"nstream/notes/updatenote"
	"nstream/status"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.SetFlags(log.Flags() + log.LUTC)
	addr := getAddr()
	log.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getAddr() string {
	host := ""
	port := os.Getenv("PORT")
	if port == "" {
		host = "localhost"
		port = "8080"
	}
	return host + ":" + port
}

func handler(response http.ResponseWriter, request *http.Request) {
	nts.Handle(request, response)
}

var nts api.Api = api.NewApi([]api.Endpoint{
	status.Status{},
	signup.Signup{},
	login.Login{},
	logout.Logout{},
	newnote.NewNote{},
	updatenote.Update{},
	deletenote.Delete{},
	fetchnotes.Fetch{},
})
