package controllers

import "github.com/gin-gonic/gin"


func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		foodId := c.Param("food_id")
		filter := bson.M{"food_id": foodId}
		var food models.Food

		err := foodCollection.FindOne(ctx, filter).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}


func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}