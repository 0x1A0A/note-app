package models

type NotesDoc struct {
	Id    string `bson:"_id"`
	Label string `bson:"label_id"`
	Name  string `bson:"name"`
	Data  string `bson:"note"`
}
