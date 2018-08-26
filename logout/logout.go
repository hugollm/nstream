package logout

import (
	"net/http"
	"nstream/auth"
	"nstream/api"
)

type Logout struct {}

func (l Logout) Accept (request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/logout"
}

func (l Logout) Handle (request *http.Request, response http.ResponseWriter) {
	_, err := auth.Authenticate(request)
	if err != nil {
		out := api.NewAuthErrorOutput()
		out.WriteToResponse(response)
		return
	}
	deleteSession(request.Header.Get("Auth-Token"))
}
