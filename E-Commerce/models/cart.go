package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Checkout  bool               `json:"checkout,omitempty" bson:"checkout"`
	// omitempty is for the values that may be come or not , so on that values we use omitempty

}