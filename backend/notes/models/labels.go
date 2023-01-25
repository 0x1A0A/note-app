package models

type LabelsDoc struct {
	Id string `bson:"_id" json:"id"` // binding:"required"
	Name string `bson:"name" json:"name"`
	User string `bson:"user" json:"user"`
}