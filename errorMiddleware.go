package HandlerFramework

import "log"

type IErrorMiddleware func(err error, req *HandlerRequest, res *HandlerResponse)

func defaultErrorMiddleware(err error, req *HandlerRequest, res *HandlerResponse) {
	log.Println("[ERROR] Calling middleware error with:", err)
	res.StatusCode = 400
	res.Message = err.Error()
}
