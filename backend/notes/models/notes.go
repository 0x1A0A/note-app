package models

type NotesDoc struct {
	Id    string `bson:"_id" json:"id"`
	Label string `bson:"label" json:"label"`
	Name  string `bson:"name" json:"name"`
	Data  string `bson:"note" json:"data"`
}
