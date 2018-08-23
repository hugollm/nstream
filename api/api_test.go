package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiCanHandleRequests(t *testing.T) {
	api := NewApi([]Endpoint{EndpointA{}})
	response := makeRequest(api, "/")
	if response.Code != 201 || response.Body.String() != "A" {
		t.Fail()
	}
}

func TestApiWritesJsonContentTypeHeader(t *testing.T) {
	api := NewApi([]Endpoint{})
	response := makeRequest(api, "/")
	if response.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestApiReturns404IfNoEndpointHandlesTheRequest(t *testing.T) {
	api := NewApi([]Endpoint{})
	response := makeRequest(api, "/foo")
	err := errors.New("Endpoint not found.")
	out := NewErrorOutput(404, map[string]error{"not_found": err})
	if response.Code != 404 || response.Body.String() != out.String() {
		t.Fail()
	}
}

func TestOnlyOneEndpointCanHandleEachRequest(t *testing.T) {
	api := NewApi([]Endpoint{
		EndpointA{},
		EndpointB{},
	})
	response := makeRequest(api, "/")
	if response.Code != 201 || response.Body.String() != "A" {
		t.Fail()
	}
}

func makeRequest(api Api, path string) *httptest.ResponseRecorder {
	request := httptest.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	api.Handle(request, response)
	return response
}

type EndpointA struct{}

func (t EndpointA) Accept(request *http.Request) bool {
	return request.URL.Path == "/"
}

func (t EndpointA) Handle(r *http.Request, w http.ResponseWriter) {
	w.WriteHeader(201)
	w.Write([]byte("A"))
}

type EndpointB struct{}

func (t EndpointB) Accept(request *http.Request) bool {
	return request.URL.Path == "/"
}

func (t EndpointB) Handle(r *http.Request, w http.ResponseWriter) {
	w.WriteHeader(202)
	w.Write([]byte("B"))
}
