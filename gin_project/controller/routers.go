package controller

import "github.com/gin-gonic/gin"

func Router(router *gin.RouterGroup){
	router.GET("/gorm_test", GormTest)
}
