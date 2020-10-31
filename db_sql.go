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
    // サーチ
	var users []relate_table.User
	db.Raw("select * from users where name = ?", 1).Find(&users)
	fmt.Println(users)

	// 増加
	//db.Exec("insert into users(name,age) values(?,?)", "hallenl",18)
	db.Exec("update users set name = ? where id = ?", "sss", 9) //update users set name = 'sss' where id = 9
}
