package main

import (
	"UserDataTestTask/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
)

type UserJSON struct {
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
	Objects []UserJSON `json:"objects"`
}

func main() {
	cl := services.ConnToMongo()

	// close connection
	defer func() {
		if err := cl.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}()

	err := MigrateUsers(cl.Database("usersdb").Collection("users"))
	if err != nil {
		log.Error(err)
	}
}

func MigrateUsers(r *mongo.Collection) error {
	//// delete all
	//del, err := r.DeleteMany(context.TODO(), bson.M{}, &options.DeleteOptions{})
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//fmt.Println("Delete: ", del.DeletedCount)

	// read .json
	data := DataJson{
		[]UserJSON{},
	}

	file, err := ioutil.ReadFile("db/migrations/users.json")
	if err != nil {
		log.Error(err)
		return err
	}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Error(err)
		return err
	}

	//INSERT ALL DOCUMENTS
	insertInterf := make([]interface{}, len(data.Objects), len(data.Objects))
	for i := range data.Objects {
		insertInterf[i] = data.Objects[i]
	}

	_, err = r.InsertMany(context.TODO(), insertInterf, &options.InsertManyOptions{})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
