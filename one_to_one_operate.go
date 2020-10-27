package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_project/models/relate_table"
)

func main(){
	// 1 on 1 の操作
	dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	d ,_ := db.DB()
	defer d.Close()

	// C
	//user := relate_table.User{
	//	Name:        "usernmae",
	//	Age:         10,
	//	Addr:        "xx",
	//	UserProfile: relate_table.UserProfile{
	//		Pic:    "1.png",
	//		CPic:   "2.png",
	//		Phone:  "1213132141",
	//	},
	//}
	//db.Create(&user)

	////// R
	// db.Model(xx).Association("xx").
	//var user relate_table.User
	//db.First(&user, 1)
	//db.Model(user).Association("UserProfile").Find(&user.UserProfile)
	//fmt.Println(user) //{1 usernmae 10 xx {1 1.png 2.png 1213132141 1}}
	//
	//// db.Preload()
	//var user2 relate_table.User
	////db.Debug().Preload("UserProfile").First(&user2, 2) //SELECT * FROM `users` WHERE `users`.`id` = 2 ORDER BY `users`.`id` LIMIT 1
	//db.Preload("UserProfile").Find(&user2, 2) //SELECT * FROM `users` WHERE `users`.`id` = 2
	//fmt.Println(user2)

	// Related 2.0以降使用できない
	//var user3 relate_table.User
	//db.First(&user3, 1)
	//fmt.Println(user3)
	//
	//var user_profile relate_table.UserProfile
	//db.Model(&user3).Related

	/// アップデート U
	//var User relate_table.User
	//db.Preload("UserProfile").First(&User, 1)
	//fmt.Println(User)
	////db.Model(&User.UserProfile).Update("Phone", "090123456")
	//db.Model(&User.UserProfile).Updates(relate_table.UserProfile{Phone:"0800808080", Pic:"dd.png"})
	//fmt.Println(User)

	///////////// D
	//// まず一つを追加
    //var user1 relate_table.User
	//db.First(&user1)
	//fmt.Println(user1)
	//db.Model(&user1).Association("UserProfile").Append(&relate_table.UserProfile{
	//	Pic:    "add.png",
	//	CPic:   "add.png",
	//	Phone:  "123",
	//})


	var user relate_table.User
	db.Preload("UserProfile").First(&user, 1)
	db.Debug().Delete(&user.UserProfile) //DELETE FROM `user_profiles` WHERE `user_profiles`.`id` = 3
}
