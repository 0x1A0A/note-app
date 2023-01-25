package controller

import (
	"0x1a0a/note-app/database"
	"0x1a0a/note-app/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func find_user(name string) *models.UsersDoc {
	col := database.DB().Collection("Users")
	res := col.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}})

	var user models.UsersDoc

	if err := res.Decode(&user); err != nil {
		return nil
	}

	return &user
}
