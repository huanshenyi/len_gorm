package main

import (
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

	// First
	//var user relate_table.User
	// 1
	//db.First(&user, 1)
	// 2
	//db.Preload("UserProfile").First(&user,"addr=?", "xxx")
	// 3
	//ret := db.Debug().Where("name = ?", "usernmae2").First(&user)
	//fmt.Println(user)
	//fmt.Println(ret.RowsAffected) // 見つけた数
	//fmt.Println(ret.Error)    // errorあるか？
	//fmt.Println(user.Id)



	// FirstOrCreate
	//var user relate_table.User
	//
	//user2 := relate_table.User{
	//	Name:"name21",
	//	Age:12,
	//	Addr:"住所1",
	//}
	//
	//ret := db.FirstOrCreate(&user, &user2)
	//fmt.Println(ret.RowsAffected)
	//fmt.Println(user)

	////// Last
	//var user6 relate_table.User
	//db.Debug().Last(&user6)
	//fmt.Println(user6)

	////// take
	//var user relate_table.User
	//db.Debug().Take(&user, 3) //SELECT * FROM `users` LIMIT 1
	//fmt.Println(user)

	////// Find
	var user []relate_table.User
	id_arry := []int{1,2,3}
	//db.Debug().Find(&user, 1) //SELECT * FROM `users` WHERE `users`.`id` = 1
	db.Debug().Find(&user, id_arry) // SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)

	//dT(id_arry...)




}

//func dT(d ...int){
//    var x  []int
//    x = append(x, d...)
//    fmt.Println(x)
//}