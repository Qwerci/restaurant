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
	Order_items	[]models.orderItem
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){

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