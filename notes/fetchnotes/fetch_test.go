package fetchnotes

import (
	"fmt"
	"net/http/httptest"
	"nstream/data/mock"
	"strings"
	"testing"
	"time"
)

var endpoint = Fetch{}

func TestAccept(t *testing.T) {
	request := httptest.NewRequest("GET", "/notes/fetch", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectMethod(t *testing.T) {
	request := httptest.NewRequest("POST", "/notes/fetch", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectPath(t *testing.T) {
	request := httptest.NewRequest("GET", "/notes/fetch/", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestAuthError(t *testing.T) {
	request := httptest.NewRequest("GET", "/notes/fetch", nil)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 401 {
		t.Fail()
	}
}

func TestJsonError(t *testing.T) {
	session := mock.Session()
	request := httptest.NewRequest("GET", "/notes/fetch", nil)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
	if !strings.Contains(response.Body.String(), "json") {
		t.Fail()
	}
}

func TestValidationError(t *testing.T) {
	session := mock.Session()
	body := strings.NewReader(`{
		"start": "0001-01-01T00:00:00Z",
		"end": "0001-01-01T00:00:00Z"
	}`)
	request := httptest.NewRequest("GET", "/notes/fetch", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
	if !strings.Contains(response.Body.String(), "start") {
		t.Fail()
	}
}

func TestFetch(t *testing.T) {
	now := time.Now()
	secondAgo := now.Add(time.Duration(-time.Second))
	yesterday := now.AddDate(0, 0, -1)
	session := mock.Session()
	note := mock.Note()
	mock.Update("notes", note.Id, "user_id", session.UserId)
	mock.Update("notes", note.Id, "created_at", secondAgo)
	sNow := now.Format(time.RFC3339)
	sYesterday := yesterday.Format(time.RFC3339)
	json := fmt.Sprintf(`{"start":"%s","end":"%s"}`, sYesterday, sNow)
	body := strings.NewReader(json)
	request := httptest.NewRequest("GET", "/notes/fetch", body)
	request.Header.Add("Auth-Token", session.Token)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 200 {
		t.Fail()
	}
	if !strings.Contains(response.Body.String(), "user_id") {
		t.Fail()
	}
}
