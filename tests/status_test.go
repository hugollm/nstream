package status

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "nts/api"
    "os"
    "testing"
)

var server *httptest.Server

func TestMain(m *testing.M) {
    server = httptest.NewServer(http.HandlerFunc(api.Handler))
    defer server.Close()
    os.Exit(m.Run())
}

func TestStatus(t *testing.T) {
    response, _ := http.Get(server.URL)
    if (response.StatusCode != 200) {
        t.Fail()
    }
}

func TestBody(t *testing.T) {
    response, _ := http.Get(server.URL)
    body, _ := ioutil.ReadAll(response.Body)
    text := string(body)
    if (text != "OK") {
        t.Fail()
    }
}
