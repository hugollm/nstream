package login

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"nstream/api"
	"nstream/data"
	"nstream/data/mock"
	"strings"
	"testing"
)

var endpoint Login = Login{}

func assertLogin(t *testing.T, input LoginInput, code int, errs map[string]error) {
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	if response.Code != code {
		t.Fail()
	}
	if len(errs) > 0 {
		expected := httptest.NewRecorder()
		api.WriteErrors(expected, 400, errs)
		if response.Code != 400 || response.Body.String() != expected.Body.String() {
			t.Fail()
		}
	}
}

func makeRequest(input LoginInput) (*http.Request, *httptest.ResponseRecorder) {
	body := strings.NewReader(mock.Json(input))
	request := httptest.NewRequest("POST", "/login", body)
	response := httptest.NewRecorder()
	return request, response
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
	user := mock.User()
	mock.Update("users", user.Id, "password", hashPassword("12345678"))
	input := LoginInput{user.Email, "12345678"}
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	token := getDbSessionTokenByUserEmail(user.Email)
	if response.Code != 200 {
		t.Fail()
	}
	if response.Body.String() != mock.Json(LoginOutput{token}) {
		t.Fail()
	}
}

func getDbSessionTokenByUserEmail(email string) (token string) {
	query := `SELECT token FROM sessions
	INNER JOIN users ON users.id = sessions.user_id
	WHERE users.email = $1 LIMIT 1`
	row := data.DB.QueryRow(query, email)
	err := row.Scan(&token)
	if err != nil {
		panic(err)
	}
	return token
}

func TestLoginWithInvalidEmail(t *testing.T) {
	input := LoginInput{"invalid-email", "12345678"}
	errs := map[string]error{"email": errors.New("Email not found.")}
	assertLogin(t, input, 400, errs)
}

func TestLoginWithInvalidPassword(t *testing.T) {
	user := mock.User()
	mock.Update("users", user.Id, "email", user.Email)
	input := LoginInput{user.Email, "invalid-password"}
	errs := map[string]error{"password": errors.New("Wrong password.")}
	assertLogin(t, input, 400, errs)
}
