package main

import (
	"UserDataTestTask/models"
	"UserDataTestTask/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
)

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
	data := models.DataJson{
		[]models.User{},
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
