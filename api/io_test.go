package api

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

type SimpleStruct struct {
	Email string `json:"email"`
}

func TestReadValidInput(t *testing.T) {
	var input SimpleStruct
	body := strings.NewReader(`{"email": "john.doe@gmail.com"}`)
	request := httptest.NewRequest("POST", "/foo", body)
	err := ReadInput(request, &input)
	if err != nil || input.Email != "john.doe@gmail.com" {
		t.Fail()
	}
}

func TestReadInvalidInput(t *testing.T) {
	var input SimpleStruct
	body := strings.NewReader("invalid-json")
	request := httptest.NewRequest("POST", "/foo", body)
	err := ReadInput(request, &input)
	if err == nil {
		t.Fail()
	}
}

func TestReadMismatchedType(t *testing.T) {
	var input SimpleStruct
	body := strings.NewReader(`"some string"`)
	request := httptest.NewRequest("POST", "/foo", body)
	err := ReadInput(request, &input)
	if err == nil {
		t.Fail()
	}
}

func TestWriteOutput(t *testing.T) {
	out := SimpleStruct{"john.doe@gmail.com"}
	response := httptest.NewRecorder()
	WriteOutput(response, 201, out)
	if response.Code != 201 {
		t.Fail()
	}
	if response.Body.String() != `{"email":"john.doe@gmail.com"}` {
		t.Fail()
	}
}

func TestWriteErrors(t *testing.T) {
	errs := map[string]error{"email": errors.New("Invalid email.")}
	response := httptest.NewRecorder()
	WriteErrors(response, 400, errs)
	if response.Code != 400 {
		t.Fail()
	}
	if response.Body.String() != `{"errors":{"email":"Invalid email."}}` {
		t.Fail()
	}
}

func TestWriteJsonError(t *testing.T) {
	response := httptest.NewRecorder()
	WriteJsonError(response)
	if response.Code != 400 {
		t.Fail()
	}
	if response.Body.String() != `{"errors":{"json":"Invalid input."}}` {
		t.Fail()
	}
}

func TestWriteAuthError(t *testing.T) {
	response := httptest.NewRecorder()
	WriteAuthError(response)
	if response.Code != 401 {
		t.Fail()
	}
	if response.Body.String() != `{"errors":{"auth":"Authentication failed."}}` {
		t.Fail()
	}
}
