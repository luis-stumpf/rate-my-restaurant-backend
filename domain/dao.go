package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"rate-my-restaurant/entity"
	"time"
)

func Create(restaurant *entity.Restaurant) (*entity.Restaurant, error) {
	restaurantC := db.Collection("restaurants")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := restaurantC.InsertOne(ctx, bson.D{
		{"name", restaurant.Name},
		{"type", restaurant.Type},
		{"rating", restaurant.Rating},
		{"image", restaurant.Image},
		{"menu", restaurant.Menu},
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	restaurant.ID = result.InsertedID.(primitive.ObjectID)
	return restaurant, nil
}

func Find(restaurantName string) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant
	restaurantC := db.Collection("restaurants")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	objID, err := primitive.ObjectIDFromHex(restaurantName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = restaurantC.FindOne(ctx, bson.M{"_id": objID}).Decode(&restaurant)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &restaurant, nil
}

func FindAll() *[]entity.Restaurant {
	var restaurants []entity.Restaurant
	restaurantC := db.Collection("restaurants")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	cursor, err := restaurantC.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var restaurant entity.Restaurant
		err := cursor.Decode(&restaurant)
		if err != nil {
			log.Fatal(err)
		}

		restaurants = append(restaurants, restaurant)
	}

	defer cursor.Close(context.Background())
	return &restaurants

}
