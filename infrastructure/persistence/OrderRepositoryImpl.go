package persistence

import (
	"context"
	"errors"
	"fmt"
	"ws_customers/domain/dto"
	utils "ws_customers/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepositoryImpl struct {
	client *mongo.Client
}

func (r *OrderRepositoryImpl) CreateOrder(o dto.OrderDto) (interface{}, error) {

	order := r.client.Database("shipping").Collection("orders")
	doc, _ := utils.ToDoc(o)
	result, err := order.InsertOne(context.TODO(), doc)

	return result.InsertedID, err
}

func (r *OrderRepositoryImpl) FindOrder(id string) dto.OrderDto {
	var result dto.OrderDto
	order := r.client.Database("shipping").Collection("orders")

	objectIDFromHex := func(hex string) primitive.ObjectID {
		objectID, err := primitive.ObjectIDFromHex(hex)
		if err != nil {
			fmt.Println(err)
		}
		return objectID
	}

	order.FindOne(context.Background(), bson.D{{"_id", objectIDFromHex(id)}}).Decode(&result)

	return result
}

func (r *OrderRepositoryImpl) ChangeStatus(idOrder string, status string) error {
	coll := r.client.Database("shipping").Collection("orders")
	id, _ := primitive.ObjectIDFromHex(idOrder)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", status}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if (result.ModifiedCount) > 0 {
		return nil
	}
	return errors.New("Not updated")
}
