package controller

import (
	"0x1a0a/note-app/database"
	"0x1a0a/note-app/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /note
func Create_note(c *gin.Context) {
	var note models.NotesDoc
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	lID, err := primitive.ObjectIDFromHex(note.Label)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("lID: %v\n", lID)

	label := find_label_by_id(lID)

	if label == nil {
		c.JSON(http.StatusBadRequest, "no label found")
		return
	}

	col := database.DB().Collection("Notes")

	res, err := col.InsertOne(context.TODO(), bson.D{
		{Key: "name", Value: note.Name},
		{Key: "label", Value: lID},
		{Key: "data", Value: note.Data},
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.InsertedID)
}

// GET notes/:label
func Get_note_by_lable_id(c *gin.Context) {
	lID, err := primitive.ObjectIDFromHex(c.Param("label"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	label := find_label_by_id(lID)

	if label == nil {
		c.JSON(http.StatusBadRequest, "no label found")
		return
	}

	col := database.DB().Collection("Notes")

	cur, err := col.Find(context.TODO(), bson.D{
		{Key: "label", Value: lID},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var res []models.NotesDoc
	if err := cur.All(context.TODO(), &res); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
