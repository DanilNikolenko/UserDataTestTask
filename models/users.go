package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id"`
	Email     string             `json:"email"`
	LastName  string             `json:"last_name"`
	Country   string             `json:"country"`
	City      string             `json:"city"`
	Gender    string             `json:"gender"`
	BirthDate string             `json:"birth_date"`
}
