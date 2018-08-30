package updatenote

import (
	"fmt"
	"net/http/httptest"
	"nstream/data/mock"
	"strings"
	"testing"
)

var endpoint = Update{}

func TestAccept(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/update", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectMethod(t *testing.T) {
	request := httptest.NewRequest("GET", "/notes/update", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectPath(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/update/", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestHandleUpdatesNote(t *testing.T) {
	session := mock.Session()
	note := mock.Note()
	mock.Update("notes", note.Id, "user_id", session.UserId)
	json := fmt.Sprintf(`{"note_id":%d,"content":"Lorem ipsum."}`, note.Id)
	body := strings.NewReader(json)
	request := httptest.NewRequest("POST", "/notes/update/", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	content := getNoteContent(note.Id)
	if response.Code != 200 {
		t.Fail()
	}
	if content != "Lorem ipsum." {
		t.Fail()
	}
}

func TestAuthError(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/update/", nil)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 401 {
		t.Fail()
	}
}

func TestJsonError(t *testing.T) {
	session := mock.Session()
	body := strings.NewReader("invalid-json")
	request := httptest.NewRequest("POST", "/notes/update/", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
}

func TestContentValidation(t *testing.T) {
	session := mock.Session()
	note := mock.Note()
	mock.Update("notes", note.Id, "user_id", session.UserId)
	json := fmt.Sprintf(`{"note_id":%d,"content":"  Lorem ipsum.  "}`, note.Id)
	body := strings.NewReader(json)
	request := httptest.NewRequest("POST", "/notes/update/", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	content := getNoteContent(note.Id)
	if response.Code != 200 {
		t.Fail()
	}
	if content != "Lorem ipsum." {
		t.Fail()
	}
}
