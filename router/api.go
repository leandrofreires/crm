package router

import "github.com/gin-gonic/gin"

//APIRouter load all routers for api
func APIRouter(r *gin.Engine) {
	userRouter(r)
	articleRouter(r)
}
