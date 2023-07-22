package handler

import "github.com/walkersumida/go-api-server/ogen"

type Handler struct {
	ogen.UnimplementedHandler
}

var _ ogen.Handler = Handler{}

func NewHandler() *Handler {
	return &Handler{}
}
