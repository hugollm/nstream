package newnote

import (
	"net/http"
	"nstream/api"
	"nstream/auth"
	"time"
)

type NewNote struct{}

type NewNoteInput struct {
	Content string
}

type NewNoteOutput struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n NewNote) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/notes/new"
}

func (n NewNote) Handle(request *http.Request, response http.ResponseWriter) {
	user, authErr := auth.Authenticate(request)
	if authErr != nil {
		api.WriteAuthError(response)
		return
	}
	var input NewNoteInput
	jsonErr := api.ReadInput(request, &input)
	if jsonErr != nil {
		api.WriteJsonError(response)
		return
	}
	vInput, _ := validateInput(input)
	note := writeNewNote(user.Id, vInput.Content)
	api.WriteOutput(response, 200, NewNoteOutput{
		Id:        note.Id,
		UserId:    note.UserId,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
	})
}
