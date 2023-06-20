package controllers

import (
	"context"
	"log"
	"net/http"
	"fmt"
	"time"
	"log"
	"github.com/Qwerci/restaurant/models"
	"github.com/Qwerci/restaurant/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu ); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		menu.Created_at = time.Now()
		menu.Updated_at = time.Now()
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, insertErr := menuCollection.InsertOne(ctx,menu)
		if insertErr != nil {
			msg := fmt.Sprintf("Menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}



func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj bson.D

		if menu.Start_Date != nil  && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()){
			msg := "kindly return time"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			defer cancel()
			return
		}

			updateObj = append(updateObj, bson.E{Key:"start_date", Value:menu.Start_Date})
			updateObj = append(updateObj, bson.E{Key:"end_date", Value:menu.End_Date})

			if menu.Name != ""{
				updateObj = append(updateObj, bson.E{Key:"name", Value:menu.Name})
			}

			if menu.Category != ""{
				updateObj = append(updateObj, bson.E{Key:"category", Value:menu.Category})
			}

			menu.Updated_at = time.Now().UTC()
			updateObj = append(updateObj, bson.E{Key:"updated_at", Value: menu.Updated_at})

			upsert := true

			opt := options.UpdateOptions{
				Upsert : &upsert,
			}

			result, err := menuCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"$set", updateObj},
				},
				&opt,
			)

			if err != nil{
				msg := "Menu update failed"
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})

			}

			defer cancel()
			c.JSON(http.StatusOK, result)
		}
	}
}