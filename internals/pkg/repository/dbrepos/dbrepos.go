package dbrepos

import (
	"context"

	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/connections"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/repository"
)

type mongoDbRepo struct {
	Client connections.DbInterface
	ctx    context.Context
}

func NewMongoDbRepo(client connections.DbInterface, ctx context.Context) repository.DbMethod {
	return &mongoDbRepo{
		Client: client,
		ctx:    ctx,
	}
}
