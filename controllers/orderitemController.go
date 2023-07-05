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

func GetOrderItemsByOrder() gin.HandlerFunc {
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

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		orderItemId := c.Param("order_item_id")
		var orderItem models.OrderItem

		err := orderitemCollection.FindOne(ctx, bson.M{"orderItem_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered item"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func ItemsByOrder( id string)(OrderItems []primitive.M, err error)


func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		
		var orderItemsPack OrderItemPack
		var order models.Order
		
		if err := c.BindJSON(&orderItemsPack); err != nil{
			c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
			return
		}

		order.Order_Date = time.Now()

		orderItemsToBeInserted := []interface{}{}
		order.Table_Order = orderItemsPack.Table_id
		order_id := OrderItemOrderCreator(order)

		for _, orderItem := range orderItemsPack.Order_items {
			orderItem.Order_id = order_id
			validationErr := validate.Struct(orderItem)

			if validationErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
				return
			}

			orderItem.ID = primitive.NewObjectID()
			orderItem.Created_at = time.Now()
			orderItem.Updated_at = time.Now()
			orderItem.Order_item_id = orderItem.ID.Hex()

			var num = toFixed(*orderItem.Unit_price, 2)
			orderItem.Unit_price = &num
			orderItemsToBeInserted = append(orderItemsToBeInserted, orderItem)
		}

		insertedOrderItems, err := orderitemCollection.InsertMany(ctx, orderItemsToBeInserted)

		if err != nil {
			log.Fatal(err)
		}
		defer cancel()

		c.JSON(http.StatusOK, insertedOrderItems)
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx,cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		var orderItem models.OrderItem
		orderItemId := c.Param("order_item_id")
		filter := bson.M{"order_item_id": orderItemId}

		var updateObj primitive.D

		if orderItem.Unit_price != nil {
			updateObj = append(updateObj, bson.E{Key:"unit_price", Value: *&orderItem.Unit_price})
		}

		if orderItem.Quantity != nil {
			updateObj = append(updateObj, bson.E{Key: "quantity", Value: *orderItem.Quantity})
		}

		if orderItem.Food_id != nil {
			updateObj = append(updateObj, bson.E{Key:"food_id", Value: *orderItem.Food_id})
		}

		orderItem.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key:"updated_at", Value: orderItem.Updated_at})

		upsert := true
		opt := options.Update().SetUpsert(upsert)

		result, err := orderitemCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			opt,
		)

		if err != nil {
			msg := "Order item update failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, result)
	}
}