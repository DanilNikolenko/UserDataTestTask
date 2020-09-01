package mongodb

import (
	"UserDataTestTask/models"
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage struct {
	Users *mongo.Collection
}

func NewMongoRepository(cl *mongo.Client, dbName, usersCollectionName string) *MongoStorage {
	rez := &MongoStorage{
		Users: cl.Database(dbName).Collection(usersCollectionName),
	}

	return rez
}

func (r *MongoStorage) GetUsersFromDB(c echo.Context) (*[]models.User, error) {
	rez := []models.User{}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	offset, err := strconv.Atoi(c.QueryParam("offset"))

	lim64 := int64(limit)
	offset64 := int64(offset)

	// ---------------Find many
	// find all canceled
	cursor, err := r.Users.Find(context.TODO(), bson.M{}, &options.FindOptions{
		Limit: &lim64,
		Skip:  &offset64,
		Sort:  bson.D{{"email", 1}},
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		err = cursor.Close(context.TODO())
		if err != nil {
			log.Error(err)
		}
	}()

	// parse all
	for cursor.Next(context.TODO()) {
		var episode models.User
		if err = cursor.Decode(&episode); err != nil {
			log.Error(err)
		}

		rez = append(rez, episode)
	}

	// ---------------Find many

	return &rez, nil
}

func (r *MongoStorage) AddUserToDB(c echo.Context, user *models.User) (*models.User, error) {

	// insert to the collection one document
	InsertedUser, err := r.Users.InsertOne(context.TODO(), user)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// RETURN USER
	rezUser := models.User{
		ID:        InsertedUser.InsertedID.(primitive.ObjectID),
		Email:     user.Email,
		LastName:  user.LastName,
		Country:   user.Country,
		City:      user.City,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}

	return &rezUser, nil
}

func (r *MongoStorage) UpdateUserInDB(c echo.Context, user *models.User) (*models.User, error) {

	// made filter for update one
	filter := bson.D{{"_id", user.ID}}

	// find user by id
	mongoUser := models.User{}
	err := r.Users.FindOne(context.TODO(), filter).Decode(&mongoUser)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if strings.TrimSpace(user.Email) != "" {
		mongoUser.Email = user.Email
	}
	if strings.TrimSpace(user.LastName) != "" {
		mongoUser.LastName = user.LastName
	}
	if strings.TrimSpace(user.Country) != "" {
		mongoUser.Country = user.Country
	}
	if strings.TrimSpace(user.City) != "" {
		mongoUser.City = user.City
	}
	if strings.TrimSpace(user.Gender) != "" {
		mongoUser.Gender = user.Gender
	}
	if strings.TrimSpace(user.BirthDate) != "" {
		mongoUser.BirthDate = user.BirthDate
	}

	// made update data for update one
	update := bson.M{"$set": bson.M{
		"email":     mongoUser.Email,
		"lastName":  mongoUser.LastName,
		"country":   mongoUser.Country,
		"city":      mongoUser.City,
		"gender":    mongoUser.Gender,
		"birthDate": mongoUser.BirthDate,
	},
	}

	_, err = r.Users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &mongoUser, nil
}
