package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"nstream/api"
	"testing"
)

var endpoint Login = Login{}

func assertLogin(t *testing.T, input LoginInput, code int, errs map[string]error) {
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	if response.Code != code {
		t.Fail()
	}
	out := api.NewErrorOutput(400, errs)
	if len(errs) > 0 && response.Body.String() != out.String() {
		t.Fail()
	}
}

func makeRequest(input LoginInput) (*http.Request, *httptest.ResponseRecorder) {
	json, _ := json.Marshal(input)
	body := bytes.NewBuffer(json)
	request := httptest.NewRequest("POST", "/login", body)
	response := httptest.NewRecorder()
	return request, response
}

func getDbSessionTokenByUserEmail(email string) (token string) {
	query := `SELECT token FROM sessions
	INNER JOIN users ON users.id = sessions.user_id
	WHERE users.email = $1 LIMIT 1`
	row := api.DB.QueryRow(query, email)
	err := row.Scan(&token)
	if err != nil {
		panic(err)
	}
	return token
}

func TestAccept(t *testing.T) {
	request := httptest.NewRequest("POST", "/login", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectMethod(t *testing.T) {
	request := httptest.NewRequest("GET", "/login", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectPath(t *testing.T) {
	request := httptest.NewRequest("POST", "/login/", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestLoginWithValidCredentials(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", hashPassword("12345678"))
	input := LoginInput{"john.doe@gmail.com", "12345678"}
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	token := getDbSessionTokenByUserEmail("john.doe@gmail.com")
	out := LoginOutput{token}
	if response.Code != 200 {
		t.Fail()
	}
	if response.Body.String() != string(out.Json()) {
		t.Fail()
	}
}

func TestLoginWithInvalidEmail(t *testing.T) {
	input := LoginInput{"invalid-email", "12345678"}
	errs := map[string]error{"email": errors.New("Email not found.")}
	assertLogin(t, input, 400, errs)
}

func TestLoginWithInvalidPassword(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", "some-hash")
	input := LoginInput{"john.doe@gmail.com", "invalid-password"}
	errs := map[string]error{"password": errors.New("Wrong password.")}
	assertLogin(t, input, 400, errs)
}
