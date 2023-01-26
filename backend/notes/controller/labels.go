package controller

import (
	"0x1a0a/note-app/database"
	"0x1a0a/note-app/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /label
func Create_label(c *gin.Context) {
	var label models.LabelsDoc

	if err := c.ShouldBindJSON(&label); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := find_user(label.User)

	if user == nil {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	col := database.DB().Collection("Labels")

	res, err := col.InsertOne(context.TODO(), bson.D{
		{Key: "name", Value: label.Name},
		{Key: "user", Value: label.User},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.InsertedID)
}

// GET /labels/:user
func Get_label_from_user(c *gin.Context) {
	user := find_user(c.Param("user"))

	if user == nil {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}
	col := database.DB().Collection("Labels")

	cur, err := col.Find(context.TODO(), bson.D{{Key: "user", Value: user.Name}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var arr []models.LabelsDoc
	if err := cur.All(context.TODO(), &arr); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, arr)
}

func find_label_by_id(id primitive.ObjectID) *models.LabelsDoc {
	col := database.DB().Collection("Labels")

	res := col.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}})

	var label models.LabelsDoc

	if err := res.Decode(&label); err != nil {
		return nil
	}

	return &label
}

// PATCH /labels/:user/:id -- TODO? don't know yet
