package repository

import (
	"context"

	"github.com/walkersumida/go-api-server/ent"
)

func (r *repository) CreateItem(ctx context.Context, data *ent.Item) (*ent.Item, error) {
	i, err := r.entcli.Item.Create().
		SetName(data.Name).
		SetStatus(data.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (r *repository) GetItems(ctx context.Context) ([]*ent.Item, error) {
	items, err := r.entcli.Item.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return items, nil
}
