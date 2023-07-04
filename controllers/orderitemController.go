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

type OrderItemPack struct{
	Table_id	*string
	Order_items	[]models.OrderItem
}

var orderitemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderitem")
func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

		result, err := orderitemCollection.Find(ctx, bson.M{})

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "error occured while getting ordered items"})
			return
		}

		var allOrderItems []bson.M
		if err = result.All(ctx, &allOrderItems); err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){
		orderId := c.Param("order_id")

		allOrderItems, err := ItemsByOrder(orderId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items by order ID"})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func ItemsByOrder( id string)(OrderItems []primitives.M, err error)


func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}