package repository

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type accountRepoositoryDB struct {
	db *mongo.Client
}

func NewAccountRepositoryDB(db *mongo.Client) accountRepoositoryDB {
	return accountRepoositoryDB{db: db}
}

func (r accountRepoositoryDB) Create(acc Account) (*Account, error) {
	insertAccount := Account{
		acc.AccountID,
		acc.CustomerID,
		acc.AccountType,
		acc.OpeningDate,
		acc.Amount,
		acc.Status,
	}
	accountColl := r.db.Database(viper.GetString("db.database")).Collection("account")
	result, err := accountColl.InsertOne(context.TODO(), insertAccount)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", result.InsertedID}}
	var accounted Account
	err = accountColl.FindOne(context.Background(), filter).Decode(&accounted)
	if err != nil {
		return nil, err
	}

	return &accounted, nil
}

func (r accountRepoositoryDB) GetAll(customerID string) ([]Account, error) {
	accountColl := r.db.Database(viper.GetString("db.database")).Collection("account")
	cursor, err := accountColl.Find(context.TODO(), bson.D{{"customer_id", customerID}})
	if err != nil {
		return nil, err
	}
	var accounts []Account
	err = cursor.All(context.TODO(), &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
