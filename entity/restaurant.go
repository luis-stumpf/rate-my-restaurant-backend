package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurant struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" binding:"required" bson:"name,omitempty"`
	Type   string             `json:"type" bson:"type,omitempty"`
	Rating float32            `json:"rating" bson:"rating,omitempty"`
	Image  string             `json:"image" bson:"image,omitempty"`
	Menu   []Dish             `json:"menu" bson:"menu"`
}

type Dish struct {
	Name  string  `json:"name" binding:"required" bson:"name,omitempty"`
	Price float32 `json:"price" binding:"required" bson:"price,omitempty"`
}

type Ingredient struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" binding:"required" bson:"name,omitempty"`
}
