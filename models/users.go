package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email"`
	LastName  string             `json:"last_name,omitempty" bson:"lastName"`
	Country   string             `json:"country,omitempty" bson:"country"`
	City      string             `json:"city,omitempty" bson:"city"`
	Gender    string             `json:"gender,omitempty" bson:"gender"`
	BirthDate string             `json:"birth_date,omitempty" bson:"birthDate"`
}

type DataJson struct {
	// !!! Necessarily use name with first BIG letter (for json marshal, unmarshal) !!!
	Objects []User `json:"objects"`
}
