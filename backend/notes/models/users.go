package models

type UsersDoc struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Password string `bson:"password"`
}