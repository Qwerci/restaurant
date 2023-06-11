package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/Qwerci/restaurant/controllers"
)


func InvoiceRoute(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices",controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id",controllers.GetInvoice())
	incomingRoutes.POST("/invoices",controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id",controllers.UpdateInvoice())
}