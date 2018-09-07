package logout

import (
	"net/http"
	"nstream/api"
	"nstream/auth"
)

type Logout struct{}

func (l Logout) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/logout"
}

func (l Logout) Handle(request *http.Request, response http.ResponseWriter) {
	_, err := auth.Authenticate(request)
	if err != nil {
		api.WriteAuthError(response)
		return
	}
	deleteSession(request.Header.Get("Auth-Token"))
}
