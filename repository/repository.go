package repository

import "github.com/walkersumida/go-api-server/ent"

type repository struct {
	entcli *ent.Client
}

func NewRepository(entcli *ent.Client) Repository {
	return &repository{entcli: entcli}
}
