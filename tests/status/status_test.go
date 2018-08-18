package status

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "nts/api"
    "testing"
)

func TestStatus(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(api.Handler))
    defer server.Close()
    response, _ := http.Get(server.URL)
    if (response.StatusCode != 200) {
        t.Fail()
    }
}

func TestBody(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(api.Handler))
    defer server.Close()
    response, _ := http.Get(server.URL)
    body, _ := ioutil.ReadAll(response.Body)
    text := string(body)
    if (text != "OK") {
        t.Fail()
    }
}
