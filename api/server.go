package api

import (
    "log"
    "net/http"
)

func RunServer() {
    http.HandleFunc("/", handleRequest)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleRequest(response http.ResponseWriter, request *http.Request) {
    for _, endpoint := range Endpoints {
        if endpoint.Accept(request) {
            endpoint.Handle(request, response)
            break
        }
    }
}
