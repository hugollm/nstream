package status

import "net/http"

type Status struct {}

func (s Status) Accept (request *http.Request) bool {
    return request.Method == "GET" && request.URL.Path == "/"
}

func (s Status) Handle (request *http.Request, response http.ResponseWriter) {
    response.Write([]byte("OK"))
}
