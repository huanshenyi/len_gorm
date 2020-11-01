package data_source

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_project/gin_project/data_source/models"
)

var Db *gorm.DB
var err error

func init(){

	mysql_conf := LoadMysqlConf()
    data_source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    	mysql_conf.UserName, mysql_conf.Password, mysql_conf.Host,mysql_conf.Port,mysql_conf.Database)

	Db, err = gorm.Open(mysql.Open(data_source), &gorm.Config{})
	if err != nil{
		panic(err.Error())
	}
	// 最大接続数を設定する
	//d,err := Db.DB()
	//d.SetMaxOpenConns(100)
	//d.SetMaxIdleConns(50)

	Db.AutoMigrate(models.GormUser{})
}
