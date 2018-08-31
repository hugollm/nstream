package deletenote

import (
	"net/http"
	"nstream/api"
	"nstream/auth"
)

type Delete struct{}

type deleteInput struct {
	NoteId int `json:"note_id"`
}

func (del Delete) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/notes/delete"
}

func (del Delete) Handle(request *http.Request, response http.ResponseWriter) {
	user, authErr := auth.Authenticate(request)
	if authErr != nil {
		api.WriteAuthError(response)
		return
	}
	var input deleteInput
	err := api.ReadInput(request, &input)
	if err != nil {
		api.WriteJsonError(response)
		return
	}
	deleteNote(user.Id, input.NoteId)
}
