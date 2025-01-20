package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name" validate:"required,min=2,max=30"`
	Email        string             `json:"email" bson:"email"`
	Phone        string             `json:"phone" bson:"phone"`
	Password     string             `json:"password" bson:"password"`
	Order_Status []Cart             `json:"order_status" bson:"order_status"`
}

type Address struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Address1 string             `json:"address" bson:"address"`
	street   string             `json:"street" bson:"street"`
	city     string             `json:"city" bson:"city"`
	Pincode  string             `json:"pincode" bson:"pincode"`
}

type Verification struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Email  string             `json:"email" bson:"email"`
	Otp    int64              `json:"otp" bson:"otp"`
	status bool               `json:"status" bson:"status"`
}
