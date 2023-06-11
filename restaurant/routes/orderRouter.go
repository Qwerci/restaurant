package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/Qwerci/restaurant/controllers"
)

func OrderRoute(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders",controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id",controllers.GetOrder())
	incomingRoutes.POST("/orders",controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id",controllers.UpdateOrder())
}