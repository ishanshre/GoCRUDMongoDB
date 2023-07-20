package repository

import (
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DbMethod interface {
	CreateProduct(p *models.CreateUpdateProduct) (*models.Product, error)
	GetProducts(limit, page int) ([]*models.Product, error)
	GetProduct(id primitive.ObjectID) (*models.Product, error)
	DeleteProduct(id primitive.ObjectID) error
	UpdateProduct(id primitive.ObjectID, update *models.CreateUpdateProduct) error
}
