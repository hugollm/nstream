package api

import "nts/common"

type Endpoint interface {
    Accept(request common.Request) bool
    Handle(request common.Request, response common.Response)
}
