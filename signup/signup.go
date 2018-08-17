package signup

import (
    "fmt"
    "net/http"
)

type Signup struct {}

func (s Signup) Accept (request *http.Request) bool {
    return request.URL.Path == "/signup"
}

func (s Signup) Handle (request *http.Request, response http.ResponseWriter) {
    fmt.Fprint(response, "Signup!")
}
