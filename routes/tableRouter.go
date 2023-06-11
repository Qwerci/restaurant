package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/Qwerci/restaurant/controllers"
)

func TableRoute(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables",controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id",controllers.GetTable())
	incomingRoutes.POST("/tables",controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id",controllers.UpdateTable())
}