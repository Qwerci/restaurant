package controllers


import (
	"context"
	"fmt"
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

type InvoiceViewFormat struct {
	Invoice_id		string
	Payment_method 	string
	Order_id		string
	Payment_status	*string
	Payment_due 	interface{}
	Table_number	interface{}
	Payment_due_date time.Time
	Order_details	interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")


func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "error occured while listing invoices"})
		}

		var allinvoices []bson.M
		if err = result.All(ctx, allinvoices); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allinvoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		invoiceId := c.Param("invoice_id")
		filter := bson.M{"invoice_id": invoiceId}
		var invoice models.Invoice

		err := invoiceCollection.FindOne(ctx, filter).Decode(&invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to list invoice"})
		}
			
	}
}


func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}