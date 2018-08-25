package signup

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"nstream/api"
	"testing"
)

var endpoint Signup = Signup{}

func TestSignup(t *testing.T) {
	tests := map[string]func(t *testing.T){
		"testAccept":                          testAccept,
		"testRejectMethod":                    testRejectMethod,
		"testRejectPath":                      testRejectPath,
		"testHandleWithProperData":            testHandleWithProperData,
		"testNewUserIsRegistered":             testNewUserIsRegistered,
		"testCannotRegisterTheSameEmailTwice": testCannotRegisterTheSameEmailTwice,
		"testInvalidJsonGetsRejected":         testInvalidJsonGetsRejected,
		"testEmailIsRequired":                 testEmailIsRequired,
		"testEmailMustBeValid":                testEmailMustBeValid,
		"testPasswordIsRequired":              testPasswordIsRequired,
		"testPasswordCanBeSpaces":             testPasswordCanBeSpaces,
		"testPasswordCannotBeTooShort":        testPasswordCannotBeTooShort,
		"testPasswordsAreHashed":              testPasswordsAreHashed,
	}
	for name, test := range tests {
		t.Run(name, test)
		clearDbUsers()
	}
}

func testAccept(t *testing.T) {
	request := httptest.NewRequest("POST", "/signup", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func testRejectMethod(t *testing.T) {
	request := httptest.NewRequest("GET", "/signup", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func testRejectPath(t *testing.T) {
	request := httptest.NewRequest("POST", "/signup/", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func testHandleWithProperData(t *testing.T) {
	input := SignupInput{"john.doe@gmail.com", "12345678"}
	assertSignup(t, input, 200, nil)
}

func testNewUserIsRegistered(t *testing.T) {
	input := SignupInput{"john.doe@gmail.com", "12345678"}
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	if !userWithEmailExists("john.doe@gmail.com") {
		t.Fail()
	}
}

func testCannotRegisterTheSameEmailTwice(t *testing.T) {
	input := SignupInput{"john.doe@gmail.com", "12345678"}
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	errors := map[string]error{"email": errors.New("Email is already taken.")}
	assertSignup(t, input, 400, errors)
}

func testInvalidJsonGetsRejected(t *testing.T) {
	body := bytes.NewBuffer([]byte("invalid-json"))
	request := httptest.NewRequest("POST", "/signup", body)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	errors := map[string]error{"json": errors.New("Invalid JSON.")}
	out := api.NewErrorOutput(400, errors)
	if response.Code != 400 || response.Body.String() != out.String() {
		t.Fail()
	}
}

func testEmailIsRequired(t *testing.T) {
	input := SignupInput{"  \n", "12345678"}
	errors := map[string]error{"email": errors.New("Email is required.")}
	assertSignup(t, input, 400, errors)
}

func testEmailMustBeValid(t *testing.T) {
	input := SignupInput{"invalid-email", "12345678"}
	errors := map[string]error{"email": errors.New("Invalid email.")}
	assertSignup(t, input, 400, errors)
}

func testPasswordIsRequired(t *testing.T) {
	input := SignupInput{"john.doe@gmail.com", ""}
	errors := map[string]error{"password": errors.New("Password is required.")}
	assertSignup(t, input, 400, errors)
}

func testPasswordCanBeSpaces(t *testing.T) {
	input := SignupInput{"john.doe@gmail.com", "        "}
	assertSignup(t, input, 200, nil)
}

func testPasswordCannotBeTooShort(t *testing.T) {
	input := SignupInput{"john.doe@gmail.com", "1234567"}
	errors := map[string]error{"password": errors.New("Password must be at least 8 characters long.")}
	assertSignup(t, input, 400, errors)
}

func testPasswordsAreHashed(t *testing.T) {
	input := SignupInput{"john.doe@gmail.com", "12345678"}
	assertSignup(t, input, 200, nil)
	password := getPasswordFromUserWithEmail("john.doe@gmail.com")
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte("12345678"))
	if err != nil {
		t.Fail()
	}
}

func getPasswordFromUserWithEmail(email string) string {
	var password string
	api.DB.QueryRow("SELECT password FROM users WHERE email = $1 LIMIT 1", email).Scan(&password)
	return password
}

func assertSignup(t *testing.T, input SignupInput, code int, errors map[string]error) {
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	if response.Code != code {
		t.Fail()
	}
	out := api.NewErrorOutput(400, errors)
	if len(errors) > 0 && response.Body.String() != out.String() {
		t.Fail()
	}
}

func makeRequest(input SignupInput) (*http.Request, *httptest.ResponseRecorder) {
	json, _ := json.Marshal(input)
	body := bytes.NewBuffer(json)
	request := httptest.NewRequest("POST", "/signup", body)
	response := httptest.NewRecorder()
	return request, response
}
