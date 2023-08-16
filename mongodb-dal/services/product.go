package services

import (
	"context"
	"fmt"
	"mongodb-dal/config"
	"mongodb-dal/models"
	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProductContext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "inventory", "products")
}
func InsertProduct(product models.Product) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := ProductContext().InsertOne(ctx, product)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}
func InsertProductList(products []interface{}) (*mongo.InsertManyResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := ProductContext().InsertMany(ctx, products)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}


