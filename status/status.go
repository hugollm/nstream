package status

import "nts/common"

type Status struct {}

func (s Status) Accept (request common.Request) bool {
    return request.Method() == "GET" && request.Path() == "/"
}

func (s Status) Handle (request common.Request, response common.Response) {
    response.WriteString("OK")
}
