package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Qwerci/restaurant/database"
	"github.com/Qwerci/restaurant/models"
	helper "github.com/Qwerci/restaurant/helpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}


func GetUser() gin.HandlerFunc {
	return func(c *gin.Context){
		
	}
}



func SignUp() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancle = context.WithTimeout(context.Background(), 100 * time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancle()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
			return
		}

		password := HashedPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exsits"})
			return
		}

		user.Created_at = time.Now().UTC()
		user.Updated_at = time.Now().UTC()
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()


		token, refreshToken, _ := helper.GenerateAllToken(*user.Email, *user.First_name, *&user.Last_name, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		//return status OK and send the result back

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}


func Login() gin.HandlerFunc {
	return func(c *gin.Context){
		
	}
}


func HashedPassword(password string) string{

}


func VerifyPassword(userPassword string, providePassword string)(bool, string){

} 