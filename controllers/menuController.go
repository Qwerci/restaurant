package controllers

import (
	"context"
	"log"
	"net/http"
	"time"
	"github.com/Qwerci/restaurant/models"
	"github.com/Qwerci/restaurant/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

		filter := bson.M{}
		results, err := menuCollection.Find(context.TODO(), filter)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured whiles listing menu items"})
		}

		var allMenus []bson.M
		if err = results.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}
		var menu models.Menu

		err := menuCollection.FindOne(ctx, filter).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu item"})
		}
		c.JSON(http.StatusOK, menu)
	}
}


func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}