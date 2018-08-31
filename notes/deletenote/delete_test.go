package deletenote

import (
	"fmt"
	"net/http/httptest"
	"nstream/data/mock"
	"strings"
	"testing"
)

var endpoint = Delete{}

func TestAccept(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/delete", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectMethod(t *testing.T) {
	request := httptest.NewRequest("GET", "/notes/delete", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectPath(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/delete/", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestNoteGetsDeleted(t *testing.T) {
	note := mock.Note()
	session := mock.Session()
	mock.Update("notes", note.Id, "user_id", session.UserId)
	json := fmt.Sprintf(`{"note_id":%d}`, note.Id)
	body := strings.NewReader(json)
	request := httptest.NewRequest("POST", "/notes/delete", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 200 {
		t.Fail()
	}
	if mock.Exists("notes", "id", note.Id) {
		t.Fail()
	}
}

func TestAuthError(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/delete", nil)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 401 {
		t.Fail()
	}
}

func TestJsonError(t *testing.T) {
	session := mock.Session()
	body := strings.NewReader("invalid-json")
	request := httptest.NewRequest("POST", "/notes/delete", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
}
