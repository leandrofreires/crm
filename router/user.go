package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leandrofreires/crm/controller"
)

func userRouter(gin *gin.Engine) {
	gin.GET("/users", controller.GetUsers)
	// gin.GET("/users/:id", controllers.GetUser)
	gin.POST("/users", controller.CreateUser)
}
