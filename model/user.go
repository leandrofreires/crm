package model

import (
	"context"
	"fmt"
	"time"

	"github.com/leandrofreires/crm/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = "user"

//User is a model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" `
	FirstName string             `bson:"first_name,omitempty" json:"first_name,omitempty" binding:"required"`
	LastName  string             `bson:"last_name,omitempty" json:"last_name,omitempty" binding:"required"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty" binding:"required,email"`
}

//FullName in database
func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

//Save user on database
func (u *User) Save() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertResult, err := database.Db.Collection(collection).InsertOne(ctx, u)
	if err != nil {
		return err
	}
	u.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

//GetUsers return of database a list of users
func (u *User) GetUsers() ([]User, error) {
	//ira receber os usuarios
	users := make([]User, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := database.Db.Collection(collection).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
