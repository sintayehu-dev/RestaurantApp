package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sintayehu-dev/go_jwt_auth/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	// Admin-only routes - restricted to restaurant staff
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())

	// Mixed access routes - permission checked inside controller
	incomingRoutes.GET("/orderItems/:order_item_id", controllers.GetOrderItem())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:order_item_id", controllers.UpdateOrderItem())
	incomingRoutes.DELETE("/orderItems/:order_item_id", controllers.DeleteOrderItem())

	// Customer-friendly routes
	incomingRoutes.GET("/orders/:order_id/items", controllers.GetOrderItemsByOrder())
}
