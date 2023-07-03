package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Qwerci/restaurant/database"
	"github.com/Qwerci/restaurant/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx,cancel = context.WithTimeout(context.Background(), 100 * time.time)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": " Failed to list tables"} )
		}

		var allTables []bson.M
		if err = result.All(ctx,&allTables); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allTables)
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx,cancel = context.WithTimeout(context.Background(), 100 * time.time)
		tableId := c.Param("table_id")
		filter := bson.M{"table_id": tableId}
		var table models.Table

		err := tableCollection.FindOne(ctx, filter).Decode(&table)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		}
		
		c.JSON(http.StatusOK, table)
	}
}


func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}