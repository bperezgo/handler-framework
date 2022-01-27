package HandlerFramework

import (
	"context"
	"errors"
)

type HandlerApp struct {
	currentIdx      int
	lengthHandlers  int
	handlers        []HandlerFunc
	ErrorMiddleware IErrorMiddleware
}

func newHandlerApp(handlers []HandlerFunc) *HandlerApp {
	return &HandlerApp{
		currentIdx:      0,
		lengthHandlers:  len(handlers),
		handlers:        handlers,
		ErrorMiddleware: defaultErrorMiddleware,
	}
}

func (ha *HandlerApp) nextFunction(req *HandlerRequest, res *HandlerResponse) NextFunction {
	return func(errs ...error) {
		if len(errs) > 0 && errs[0] != nil {
			// Call the middleware error with the first error
			ha.ErrorMiddleware(errs[0], req, res)
			return
		}
		// Call the next handler
		if ha.currentIdx == ha.lengthHandlers-1 {
			// Call the middleware error informing, that is the last handler
			// And it has not returned a response
			err := errors.New("[ERROR] Last handler")
			ha.ErrorMiddleware(err, req, res)
			return
		}
		// Call the next handler
		ha.currentIdx += 1
		ha.callHandler(ha.handlers[ha.currentIdx], req, res, req.next)
	}
}

func (ha *HandlerApp) callHandler(handler HandlerFunc, req *HandlerRequest, res *HandlerResponse, next NextFunction) {
	handler(req, res, next)
}

func (ha *HandlerApp) Handle(ctx context.Context, req interface{}) (res HandlerResponse) {
	handlerResponse := &HandlerResponse{}
	handlerReq := &HandlerRequest{
		Body: req,
	}
	nextFunction := ha.nextFunction(handlerReq, handlerResponse)
	handlerReq.next = nextFunction
	ha.callHandler(ha.handlers[0], handlerReq, handlerResponse, nextFunction)
	return *handlerResponse
}
