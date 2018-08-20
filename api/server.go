package api

import (
    "errors"
    "log"
    "net/http"
    "nts/common"
)

func RunServer() {
    http.HandleFunc("/", Handler)
    log.Println("Starting server on localhost:8080")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func Handler(response http.ResponseWriter, request *http.Request) {
    found := false
    for _, endpoint := range Endpoints {
        if endpoint.Accept(request) {
            endpoint.Handle(request, response)
            found = true
            break
        }
    }
    if !found {
        err := errors.New("Specified endpoint is unknown.")
        out := common.NewErrorOutput(404, map[string]error{"not-found": err})
        out.WriteToResponse(response)
    }
}
