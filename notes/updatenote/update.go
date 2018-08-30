package updatenote

import (
	"net/http"
	"nstream/api"
	"nstream/auth"
)

type Update struct{}

type updateInput struct {
	NoteId  int    `json:"note_id"`
	Content string `json:"content"`
}

func (up Update) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/notes/update"
}

func (up Update) Handle(request *http.Request, response http.ResponseWriter) {
	user, authErr := auth.Authenticate(request)
	if authErr != nil {
		api.WriteAuthError(response)
		return
	}
	var input updateInput
	err := api.ReadInput(request, &input)
	if err != nil {
		api.WriteJsonError(response)
		return
	}
	vInput, _ := validateInput(input)
	updateNote(user.Id, vInput.NoteId, vInput.Content)
}
