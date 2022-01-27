package HandlerFramework

import (
	"context"
	"errors"
	"log"
	"reflect"
	"testing"
)

func TestOkChainedHandlers(t *testing.T) {
	testName := "should execute all the handlers of the chain when the chain is ok"
	type body struct {
		value string
	}
	reqValue := HandlerRequest{
		Body: body{value: "value"},
	}
	resValue := HandlerResponse{
		StatusCode: 200,
	}
	handler1 := func(req *HandlerRequest, res *HandlerResponse, next NextFunction) {
		log.Println("[INFO] First call doing something")
		next()
	}
	handler2 := func(req *HandlerRequest, res *HandlerResponse, next NextFunction) {
		log.Println("[INFO] Second call doing something")
		next()
	}
	handler3 := func(req *HandlerRequest, res *HandlerResponse, next NextFunction) {
		log.Println("[INFO] Third call changing the response")
		res.StatusCode = 200
	}
	mainHandler := ComposeHandlers(handler1, handler2, handler3)
	t.Run(testName, func(t *testing.T) {
		log.Println("[INFO] TEST NAME:", testName)
		ctx := context.Background()
		resGotten := mainHandler.Handle(ctx, reqValue)
		if !reflect.DeepEqual(resGotten, resValue) {
			t.Errorf("The response gotten %+v, is not equal to response expected %+v", resGotten, resValue)
		}
	})
}

func TestChainWithThreeHandlersAndInTheSecondGetError(t *testing.T) {
	testName := "should return error in the second handler, and call the error middleware"
	type body struct {
		value string
	}
	reqValue := HandlerRequest{
		Body: body{value: "value"},
	}
	resValue := HandlerResponse{
		StatusCode: 400,
		Message:    "Some error ocurred",
	}
	handler1 := func(req *HandlerRequest, res *HandlerResponse, next NextFunction) {
		log.Println("[INFO] First call doing something")
		res.StatusCode = 200
		next()
	}
	handler2 := func(req *HandlerRequest, res *HandlerResponse, next NextFunction) {
		log.Println("[INFO] Second call doing something")
		res.StatusCode = 201
		err := errors.New("Some error ocurred")
		next(err)
	}
	handler3 := func(req *HandlerRequest, res *HandlerResponse, next NextFunction) {
		log.Println("[INFO] Third call changing the response, This log wont be printed")
		res.StatusCode = 202
	}
	mainHandler := ComposeHandlers(handler1, handler2, handler3)
	t.Run(testName, func(t *testing.T) {
		log.Println("[INFO] TEST NAME:", testName)
		ctx := context.Background()
		resGotten := mainHandler.Handle(ctx, reqValue)
		if !reflect.DeepEqual(resGotten, resValue) {
			t.Errorf("The response gotten %+v, is not equal to response expected %+v", resGotten, resValue)
		}
	})
}
