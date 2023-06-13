package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Signup struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Name  string             `json:"category,omitempty" bson:"category,omitempty"`
	Password string      `json:"password,omitempty" bson:"password,omitempty"`
	// Image []byte           `json:"website,omitempty" bson:"website,omitempty"`
}
