package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leandrofreires/crm/controller"
)

func articleRouter(gin *gin.Engine) {
	gin.GET("/articles", controller.GetArticles)
	gin.GET("/articles/:id", controller.GetArticle)
	gin.POST("/articles", controller.CreateArticle)
	gin.PUT("/articles/:id", controller.UpdateArticle)
}
