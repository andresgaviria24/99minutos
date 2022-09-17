package persistence

import (
	"ws_customers/infrastructure/repository"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbHelper struct {
	OrderRepository repository.OrderRepository
	db              *mongo.Client
}

func InitDbHelper() (*DbHelper, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	return &DbHelper{
		OrderRepository: &OrderRepositoryImpl{client},
		db:              client,
	}, nil
}
