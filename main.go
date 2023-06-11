package main

import(
	"os"
	"github.com/gin-gonic/gin"
	"github.com/Qwerci/restaurant/routes"
	"github.com/Qwerci/restaurant/database"
	"github.com/Qwerci/restaurant/middlewares"

)


func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)

}