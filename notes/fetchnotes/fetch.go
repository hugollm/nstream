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
	input, inputErr := readInput(request)
	if inputErr != nil {
		api.WriteJsonError(response)
		return
	}
	errs := validateInput(input)
	if len(errs) > 0 {
		api.WriteErrors(response, 400, errs)
		return
	}
	notes := fetchNotes(user.Id, input.Start, input.End)
	jOut := jsonOutput{notes}
	api.WriteOutput(response, 200, jOut)
}

func readInput(request *http.Request) (fetchInput, error) {
	input := fetchInput{}
	jInput := jsonInput{}
	jsonErr := api.ReadInput(request, &jInput)
	if jsonErr != nil {
		return input, jsonErr
	}
	start, startErr := time.Parse(time.RFC3339, jInput.Start)
	if startErr != nil {
		println(startErr.Error())
		return input, startErr
	}
	end, endErr := time.Parse(time.RFC3339, jInput.End)
	if endErr != nil {
		return input, endErr
	}
	input.Start = start
	input.End = end
	return input, nil
}
