package handler

import (
	"github.com/walkersumida/go-api-server/ogen"
	"github.com/walkersumida/go-api-server/repository"
)

type Handler struct {
	ogen.UnimplementedHandler

	repo repository.Repository
}

var _ ogen.Handler = Handler{}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}
