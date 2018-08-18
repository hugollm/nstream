package status

import (
    "net/http/httptest"
    "testing"
)

var endpoint Status = Status{}

func TestAccept(t *testing.T) {
    request := httptest.NewRequest("GET", "/", nil)
    if !endpoint.Accept(request) {
        t.Fail()
    }
}

func TestReject(t *testing.T) {
    request := httptest.NewRequest("GET", "/foo", nil)
    if endpoint.Accept(request) {
        t.Fail()
    }
}

func TestHandle(t *testing.T) {
    request := httptest.NewRequest("GET", "/", nil)
    response := httptest.NewRecorder()
    endpoint.Handle(request, response)
    if (response.Code != 200) {
        t.Fail()
    }
    if (response.Body.String() != "OK") {
        t.Fail()
    }
}
