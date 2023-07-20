package dbrepos

import (
	"context"
	"errors"

	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDbRepo) CreateProduct(p *models.CreateUpdateProduct) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()

	res, err := m.Client.GetProductCollection().InsertOne(ctx, p)
	if err != nil {
		return nil, errors.New("error in inserting product")
	}
	prodct, _ := m.GetProduct(res.InsertedID.(primitive.ObjectID))
	return prodct, nil
}

func (m *mongoDbRepo) GetProducts(limit, page int) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()

	if limit == 0 || limit < 0 {
		limit = 10
	}
	if page == 0 || page < 0 {
		page = 1
	}
	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	query := bson.M{}

	res, err := m.Client.GetProductCollection().Find(ctx, query, &opt)
	if err != nil {
		return nil, errors.New("error fetching all products")
	}
	products := []*models.Product{}
	for res.Next(ctx) {
		product := &models.Product{}
		if err := res.Decode(&product); err != nil {
			return nil, errors.New("error in scanning product")
		}
		products = append(products, product)
	}
	return products, nil
}

func (m *mongoDbRepo) GetProduct(id primitive.ObjectID) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()

	res := m.Client.GetProductCollection().FindOne(ctx, bson.M{"_id": id})
	product := &models.Product{}
	if err := res.Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("product not found")
		}
		return nil, errors.New("error in fetching product")
	}
	return product, nil
}

func (m *mongoDbRepo) DeleteProduct(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()

	res, err := m.Client.GetProductCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.New("error in deleteing product")
	}
	if res.DeletedCount == 0 {
		return errors.New("no product deleted")
	}
	return nil
}

func (m *mongoDbRepo) UpdateProduct(id primitive.ObjectID, update *models.CreateUpdateProduct) error {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()
	updateQuery := bson.D{{"$set", update}}
	_, err := m.Client.GetProductCollection().UpdateOne(ctx, bson.M{"_id": id}, updateQuery)
	if err != nil {
		return err
	}

	return nil

}
