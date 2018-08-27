package new

import (
	"net/http/httptest"
	"nstream/data/mock"
	"strconv"
	"strings"
	"testing"
)

var endpoint NewNote = NewNote{}

func TestAccept(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/new", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectMethod(t *testing.T) {
	request := httptest.NewRequest("GET", "/notes/new", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectPath(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/new/", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestAuthError(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/new/", nil)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 401 {
		t.Fail()
	}
}

func TestJsonError(t *testing.T) {
	session := mock.Session()
	request := httptest.NewRequest("POST", "/notes/new/", nil)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
}

func TestSuccessHandle(t *testing.T) {
	session := mock.Session()
	input := NewNoteInput{"  Lorem ipsum.  \n"}
	body := strings.NewReader(mock.Json(input))
	request := httptest.NewRequest("POST", "/notes/new/", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 200 {
		t.Fail()
	}
	if !mock.Exists("notes", "user_id", session.UserId) {
		t.Fail()
	}
}

func TestSuccessOutput(t *testing.T) {
	session := mock.Session()
	input := NewNoteInput{"  Lorem ipsum.  \n"}
	body := strings.NewReader(mock.Json(input))
	request := httptest.NewRequest("POST", "/notes/new/", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	out := response.Body.String()
	if !strings.Contains(out, `"id"`) {
		t.Fail()
	}
	if !strings.Contains(out, `"user_id":`+strconv.Itoa(session.UserId)) {
		t.Fail()
	}
	if !strings.Contains(out, `"content":"Lorem ipsum."`) {
		t.Fail()
	}
	if !strings.Contains(out, `"created_at"`) {
		t.Fail()
	}
}
