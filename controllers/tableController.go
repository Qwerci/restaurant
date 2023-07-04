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
		var ctx,cancel = context.WithTimeout(context.Background(), 100 * time.Second)

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
		var ctx,cancel = context.WithTimeout(context.Background(), 100 * time.Second)
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
		var ctx,cancel = context.WithTimeout(context.Background(), 100 * time.Second)

		var table models.Table

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(table)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		table.Created_at = time.Now()
		table.Updated_at = time.Now()

		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()

		result, insertErr := tableCollection.InsertOne(ctx, table)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Table was not created"})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx,cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		tableId := c.Param("table_id")
		var table models.Table

		var updateObj primitive.D

		if table.Number_of_guests != nil {
			updateObj = append(updateObj, bson.E{Key: "number_of_guests", Value: table.Number_of_guests})
		}

		if table.Table_number != nil {
			updateObj = append(updateObj, bson.E{Key: "table_number", Value: table.Table_number})
		}

		table.Updated_at = time.Now()

		upsert := true
		opt := options.Update().SetUpsert(upsert)

		filter := bson.M{"table_id": tableId}

		result, err :=tableCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key:"$set", Value: updateObj},
			},
			opt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, 
				gin.H{"error": "table item update failed"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}