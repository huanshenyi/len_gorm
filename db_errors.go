package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_project/models/relate_table"
)

func main(){
	dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	var users []relate_table.User
	result := db.Where("name = ?", "name2").Find(&users)
	if result.Error != nil{
      fmt.Println(result.Error.Error())
	}
	fmt.Println(result.RowsAffected)
	fmt.Println(users)


	ct := db.Begin()
	ret := ct.Commit()
	if ret.Error != nil{
		ct.Rollback()
	}

}
