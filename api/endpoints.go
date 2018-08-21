package api

import (
    "nstream/signup"
    "nstream/status"
)

var Endpoints []Endpoint = []Endpoint {
    status.Status{},
    signup.Signup{},
}
