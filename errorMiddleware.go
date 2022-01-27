package HandlerFramework

import "log"

type IErrorMiddleware func(err error, req interface{}, res interface{})

func defaultErrorMiddleware(err error, req interface{}, res interface{}) {
	log.Println("[ERROR] Calling middleware error with:", err)
}
