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
var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")
var ctx, cancel = context.WithTimeout(context.Background, 100 * time.Second)


func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context){

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

		var order models.Order
		var table models.Table

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		}

		validatorErr := validate.Struct(order)
		if validatorErr != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": validatorErr.Error()})
		}

		if order.Table_id != nil{
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}

		order.Created_at = time.Now()
		order.Updated_at = time.Now()

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": " order item was not created"})
		}

		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}