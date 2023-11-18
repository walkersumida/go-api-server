package handler

import (
	"context"

	"github.com/walkersumida/go-api-server/ent"
	"github.com/walkersumida/go-api-server/ent/item"
	"github.com/walkersumida/go-api-server/ogen"
)

func (h Handler) CreateItem(ctx context.Context, req *ogen.Item) (*ogen.Item, error) {
	item := &ent.Item{
		Name:   req.Name,
		Status: item.Status(req.Status),
	}
	i, err := h.repo.CreateItem(ctx, item)
	if err != nil {
		return nil, err
	}

	return &ogen.Item{
		ID:     ogen.NewOptUUID(i.ID),
		Name:   i.Name,
		Status: ogen.ItemStatus(i.Status),
	}, nil
}

func (h Handler) GetItems(ctx context.Context) ([]ogen.Item, error) {
	items, err := h.repo.GetItems(ctx)
	if err != nil {
		return nil, err
	}

	var res []ogen.Item
	for _, item := range items {
		res = append(res, ogen.Item{
			ID:     ogen.NewOptUUID(item.ID),
			Name:   item.Name,
			Status: ogen.ItemStatus(item.Status),
		})
	}

	return res, nil
}
