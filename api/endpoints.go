package api

import "nts/signup"

var Endpoints []Endpoint = []Endpoint {
    signup.Signup{},
}
