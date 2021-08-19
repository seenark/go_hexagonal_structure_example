package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepositoryDB struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCustomerRepositoryDB(collection *mongo.Collection, ctx context.Context) CustomerRepository {
	return customerRepositoryDB{
		collection: collection,
		ctx:        ctx,
	}
}

func (c customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	cur, err := c.collection.Find(c.ctx, bson.M{})
	if err != nil {
		return customers, err
	}
	for cur.Next(c.ctx) {
		customer := Customer{}
		cur.Decode(&customer)
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c customerRepositoryDB) GetById(id string) (*Customer, error) {
	customer := Customer{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &customer, err
	}
	err = c.collection.FindOne(c.ctx, bson.M{"_id": _id}).Decode(&customer)
	if err != nil {
		return &customer, err
	}
	return &customer, nil
}
