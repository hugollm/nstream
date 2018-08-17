package api

import (
    "log"
    "net/http"
    "nts/common"
)

func RunServer() {
    http.HandleFunc("/", handleRequest)
    log.Println("Starting server on localhost:8080")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    request := common.Request{r}
    response := common.Response{w}
    for _, endpoint := range Endpoints {
        if endpoint.Accept(request) {
            endpoint.Handle(request, response)
            break
        }
    }
}
