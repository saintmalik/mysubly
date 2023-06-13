package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"name,omitempty" bson:"name,omitempty"`
	Name  string             `json:"category,omitempty" bson:"category,omitempty"`
	Image []byte           `json:"website,omitempty" bson:"website,omitempty"`
}
