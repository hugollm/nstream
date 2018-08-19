package common

import (
    "encoding/json"
    "net/http"
)

type ApiError struct {
    Code int `json:"code"`
    Key string `json:"key"`
    Message string `json:"message"`
}

func NewApiError(code int, key string, message string) ApiError {
    return ApiError{code, key, message}
}

func NewJsonError() ApiError {
    return NewApiError(400, "invalid-json", "Received input is not valid JSON.")
}

func (err ApiError) Error() string {
    return string(err.Json())
}

func (apiErr ApiError) Json() []byte {
    dict := make(map[string]ApiError)
    dict["error"] = apiErr
    json, err := json.Marshal(dict)
    if err != nil {
        panic(err)
    }
    return json
}

func (err ApiError) WriteToResponse(response http.ResponseWriter) {
    response.WriteHeader(err.Code)
    response.Write(err.Json())
}
