package api

import (
    "nts/signup"
    "nts/status"
)

var Endpoints []Endpoint = []Endpoint {
    status.Status{},
    signup.Signup{},
}
