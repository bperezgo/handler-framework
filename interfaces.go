package HandlerFramework

import "context"

type HandlerRequest struct {
	Body interface{}
	next NextFunction
}

type HandlerResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type NextFunction func(err ...error)
type HandlerFunc func(req *HandlerRequest, res *HandlerResponse, next NextFunction)
type IHandler interface {
	Handle(ctx context.Context, req interface{}) (res interface{}, err error)
}
