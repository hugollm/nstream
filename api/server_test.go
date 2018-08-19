package api

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "nts/common"
    "os"
    "testing"
)

var server *httptest.Server

func TestMain(m *testing.M) {
    server = httptest.NewServer(http.HandlerFunc(Handler))
    defer server.Close()
    os.Exit(m.Run())
}

func TestStatus(t *testing.T) {
    response, _ := http.Get(server.URL)
    body, _ := ioutil.ReadAll(response.Body)
    text := string(body)
    if (response.StatusCode != 200 || text != "OK") {
        t.Fail()
    }
}

func TestNotFound(t *testing.T) {
    response, _ := http.Get(server.URL + "/not-found")
    body, _ := ioutil.ReadAll(response.Body)
    text := string(body)
    expectedBody := common.NewNotFoundError().Error()
    if (response.StatusCode != 404 || text != expectedBody) {
        t.Fail()
    }
}
