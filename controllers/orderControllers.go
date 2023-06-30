package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Qwerci/restaurant/database"
	"github.com/Qwerci/restaurant/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background, 100 * time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "error occured while listing order items"})	
		}
		var allOrders []bson.M{}
		if err = result.All(ctx, &allOrders); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
	}

}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background, 100 * time.Second)
		
		var order models.Order
		orderId := c.Param("order_id")
		filter := bson.M{"order_id": orderId}

		err := menuCollection.FindOne(ctx, filter).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, 
			gin.H{"error": "error occured while fetching the order"})
		}
		c.Json(http.StatusOK, order)
	}
}


func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}