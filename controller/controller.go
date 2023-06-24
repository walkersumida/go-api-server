package controller

import "github.com/walkersumida/go-api-server/ogen"

type Controller struct {
	ogen.UnimplementedHandler
}

var _ ogen.Handler = Controller{}

func NewController() *Controller {
	return &Controller{}
}
