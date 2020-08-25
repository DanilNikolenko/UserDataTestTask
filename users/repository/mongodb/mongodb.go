package mongodb

import (
	"UserDataTestTask/models"
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	if err != nil {
		log.Error(err)
		return nil, err
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		log.Error(err)
		return nil, err
	}

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
		var episode UserJSON
		if err = cursor.Decode(&episode); err != nil {
			log.Error(err)
		}
		temp := models.User{
			ID:        episode.ID,
			Email:     episode.Email,
			LastName:  episode.LastName,
			Country:   episode.Country,
			City:      episode.City,
			Gender:    episode.Gender,
			BirthDate: episode.BirthDate,
		}
		rez = append(rez, temp)
	}

	// ---------------Find many

	return &rez, nil
}

func (r *MongoStorage) AddUserToDB(c echo.Context) (*models.User, error) {
	u := UserJSON{}

	// Decode request
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// close BODY req
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	// insert to the collection one document
	rezInsert, err := r.Users.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// RETURN USER
	rezUser := models.User{
		ID:        rezInsert.InsertedID.(primitive.ObjectID),
		Email:     u.Email,
		LastName:  u.LastName,
		Country:   u.Country,
		City:      u.City,
		Gender:    u.Gender,
		BirthDate: u.BirthDate,
	}

	return &rezUser, nil
}

func (r *MongoStorage) UpdateUserInDB(c echo.Context) (*models.User, error) {
	user := UserJSON{}

	// decode request
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// close BODY req
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	// made filter for update one
	filter := bson.D{{"_id", user.ID}}

	// find user by id
	mongoUser := UserJSON{}
	err = r.Users.FindOne(context.TODO(), filter).Decode(&mongoUser)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if user.Email != "" {
		mongoUser.Email = user.Email
	}
	if user.LastName != "" {
		mongoUser.LastName = user.LastName
	}
	if user.Country != "" {
		mongoUser.Country = user.Country
	}
	if user.City != "" {
		mongoUser.City = user.City
	}
	if user.Gender != "" {
		mongoUser.Gender = user.Gender
	}
	if user.BirthDate != "" {
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

	// RETURN USER
	rezUser := models.User{
		ID:        mongoUser.ID,
		Email:     mongoUser.Email,
		LastName:  mongoUser.LastName,
		Country:   mongoUser.Country,
		City:      mongoUser.City,
		Gender:    mongoUser.Gender,
		BirthDate: mongoUser.BirthDate,
	}

	return &rezUser, nil
}