package controller

import (
	"context"

	"github.com/walkersumida/go-api-server/ogen"
)

func(c Controller) AddItem(ctx context.Context, req *ogen.Item) (*ogen.Item, error) {
	return &ogen.Item{}, nil
}

func (c Controller) GetItem(ctx context.Context, params ogen.GetItemParams) (*ogen.Item, error) {
	return &ogen.Item{}, nil
}
