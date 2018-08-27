package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ReadInput(request *http.Request, input interface{}) error {
	return json.NewDecoder(request.Body).Decode(input)
}

func WriteOutput(response http.ResponseWriter, status int, output interface{}) {
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	response.WriteHeader(status)
	response.Write(jsonBytes)
}

func WriteErrors(response http.ResponseWriter, status int, errs map[string]error) {
	output := make(map[string]map[string]string)
	output["errors"] = make(map[string]string)
	for key, value := range errs {
		output["errors"][key] = value.Error()
	}
	WriteOutput(response, status, output)
}

func WriteJsonError(response http.ResponseWriter) {
	errs := map[string]error{"json": errors.New("Invalid input.")}
	WriteErrors(response, 400, errs)
}

func WriteAuthError(response http.ResponseWriter) {
	errs := map[string]error{"auth": errors.New("Authentication failed.")}
	WriteErrors(response, 401, errs)
}
