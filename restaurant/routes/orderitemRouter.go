package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/Qwerci/restaurant/controllers"
)

func OrderItemRoute(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItems",controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id",controllers.GetOrderItem())
	incomingRoutes.POST("/orderItems",controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id",controllers.UpdateOrderItem())
}