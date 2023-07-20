package handlers

import "github.com/ishanshre/GoCRUDMongoDB/internals/pkg/repository"

type Handlers interface{}

type handler struct {
	MG repository.DbMethod
}

func NewHandler(mg repository.DbMethod) Handlers {
	return &handler{
		MG: mg,
	}
}
