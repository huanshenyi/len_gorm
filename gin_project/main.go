package main

import (
	"github.com/gin-gonic/gin"
	_ "gorm_project/gin_project/data_source"
	"gorm_project/gin_project/router"
)

func main(){
	engin := gin.Default()
    router.Router(engin)
	engin.Run()
}
