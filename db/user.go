package db

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/agchavez/tongo-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindUserByEmail(email string) (bool, models.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := MongoCon.Database("tongo").Collection("users")
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	//11ID := user.ID.Hex()
	if err != nil {
		return false, user
	}
	return true, user
}

func SaveNewUser(user models.User) (models.User, bool) {
	log.Println(user)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := MongoCon.Database("tongo").Collection("users")
	user.Password, _ = EncripPassword(user.Password)
	user.CreatedAt = time.Now()
	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		log.Fatal(err.Error())
		return user, true
	}
	return user, false
}

// Get all users
func GetUsers(
	limit string,
	offset string,
) []models.User {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	opts := options.Find()
	if len(limit) > 0 {
		limitInt, _ := strconv.ParseInt(limit, 10, 64)
		opts = opts.SetLimit(limitInt)
	}
	if len(offset) > 0 {
		offsetInt, _ := strconv.ParseInt(offset, 10, 64)
		opts = opts.SetSkip(offsetInt)
	}
	collection := MongoCon.Database("tongo").Collection("users")
	var users []models.User
	cur, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Println("Error al obtener los usuarios 1")
	}
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Println("Error al obtener los usuarios 2")
		}
		users = append(users, user)
	}
	return users
}
