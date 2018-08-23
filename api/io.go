package api

import (
	"encoding/json"
	"net/http"
)

type ErrorOutput struct {
	code   int
	errors map[string]error
}

func NewErrorOutput(code int, errors map[string]error) ErrorOutput {
	return ErrorOutput{code, errors}
}

func (out ErrorOutput) WriteToResponse(response http.ResponseWriter) {
	response.WriteHeader(out.code)
	response.Write(out.Json())
}

func (out ErrorOutput) Json() []byte {
	data := make(map[string]map[string]string)
	data["errors"] = make(map[string]string)
	for key, value := range out.errors {
		data["errors"][key] = value.Error()
	}
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return json
}

func (out ErrorOutput) String() string {
	return string(out.Json())
}
