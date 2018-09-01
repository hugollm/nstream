package fetchnotes

import (
	"net/http"
	"nstream/api"
	"nstream/auth"
	"nstream/data"
	"time"
)

type Fetch struct{}

type jsonInput struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type fetchInput struct {
	Start time.Time
	End   time.Time
}

type jsonOutput struct {
	Notes []data.Note `json:"notes"`
}

func (ft Fetch) Accept(request *http.Request) bool {
	return request.Method == "GET" && request.URL.Path == "/notes/fetch"
}

func (ft Fetch) Handle(request *http.Request, response http.ResponseWriter) {
	user, authErr := auth.Authenticate(request)
	if authErr != nil {
		api.WriteAuthError(response)
		return
	}
	jsInput := jsonInput{}
	jsonErr := api.ReadInput(request, &jsInput)
	if jsonErr != nil {
		api.WriteJsonError(response)
		return
	}
	vInput, errs := validateInput(jsInput)
	if len(errs) > 0 {
		api.WriteErrors(response, 400, errs)
		return
	}
	notes := fetchNotes(user.Id, vInput.Start, vInput.End)
	jOut := jsonOutput{notes}
	api.WriteOutput(response, 200, jOut)
}
