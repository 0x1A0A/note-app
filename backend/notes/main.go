package main

import (
	"0x1a0a/note-app/controller"
	"0x1a0a/note-app/database"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, "+
				"Content-Length, "+
				"Accept-Encoding, "+
				"X-CSRF-Token, "+
				"Authorization, "+
				"accept, "+
				"origin, "+
				"Cache-Control, "+
				"X-Requested-With",
		)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	database.Connect()

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/label", controller.Create_label)
	router.GET("/labels/:user", controller.Get_label_from_user)

	router.Run("localhost:6000")

	defer func() {
		ctx, cancle := context.WithTimeout(context.Background(), time.Second*2)
		defer cancle()
		if err := database.DB().Client().Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
