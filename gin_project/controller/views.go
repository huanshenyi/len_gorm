package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm_project/gin_project/data_source"
	"gorm_project/gin_project/data_source/models"
	"net/http"
)

func GormTest(ctx *gin.Context){
	user := models.GormUser{
		Name: "test",
		Age:  1,
	}
	ret := data_source.Db.Create(&user)
	if ret.Error != nil{
		fmt.Println("error!!!!!!:", ret.Error.Error())
	}
	ctx.String(http.StatusOK, "ok")
}
