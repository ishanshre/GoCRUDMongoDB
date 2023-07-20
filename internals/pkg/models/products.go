package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       int                `json:"price,omitempty" bson:"price,omitempty"`
	Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
}

type CreateUpdateProduct struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Price       int    `json:"price,omitempty" bson:"price,omitempty"`
	Stock       int    `json:"stock,omitempty" bson:"stock,omitempty"`
}
