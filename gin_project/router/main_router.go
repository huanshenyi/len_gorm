package router

import (
	"github.com/gin-gonic/gin"
	"gorm_project/gin_project/controller"
)

func Router(router *gin.Engine){
	chap := router.Group("/chap")
	controller.Router(chap)
}
