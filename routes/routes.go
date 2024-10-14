package routes

import (
	"fitpass/controller"

	"github.com/gin-gonic/gin"
)


func AllRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("get-user", controller.GetUsers)
	incomingRoutes.POST("add-user", controller.AddUser)
	incomingRoutes.DELETE("delete-user", controller.DeleteUser)
	incomingRoutes.POST("subscribe", controller.AddUserPlan)
}