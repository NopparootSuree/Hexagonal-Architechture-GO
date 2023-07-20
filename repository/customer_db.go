package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepositoryDB struct {
	db *mongo.Client
}

func NewCustomerRepositoryDB(db *mongo.Client) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	cutomerColl := r.db.Database("company").Collection("customer")
	cursor, err := cutomerColl.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var customers []Customer
	err = cursor.All(context.TODO(), &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r customerRepositoryDB) GetOne(name string) (*Customer, error) {
	cutomerColl := r.db.Database("company").Collection("customer")
	var customer Customer

	err := cutomerColl.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}

	return &customer, nil
}
