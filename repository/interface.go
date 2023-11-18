package repository

import (
	"context"

	"github.com/walkersumida/go-api-server/ent"
)

type Repository interface {
	CreateItem(ctx context.Context, data *ent.Item) (*ent.Item, error)
	GetItems(ctx context.Context) ([]*ent.Item, error)
}
