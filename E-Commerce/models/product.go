package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Product struct {
	ID          primitive.ObjectID     `json:"id" bson:"_id"`
	Name        string                 `json:"name" bson:"name"`
	Description string                 `json:"description" bson:"description"`
	Price       float64                `json:"price" bson:"price"`
	Image_url   string                 `json:"image_url" bson:"image_url"`
	Metainfo    map[string]interface{} `json:"meta_info,omitempty" bson:"meta_info,omitempty"` // Metainfo is to add any extra detail regrading the product
	// omitempty is for the values may be come or not , so on that values we use omitempty
}