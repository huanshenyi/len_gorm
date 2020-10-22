package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_project/models"
)

func main () {
   dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"	
   v2_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
   if err != nil{
   	  panic(err.Error())
   }
   db, _ :=v2_db.DB()
   defer db.Close()

   v2_db.AutoMigrate(&models.User{}, &models.UserInfo{}, models.Article{})
   //v2_db.AutoMigrate(&models.GormModel{})

   // C
   //v2_db.Create(&models.User{
	//   Name:  "name",
	//   Age:   10,
	//   Addr:  "xxx",
	//   Pic:   "/static/upload/pic111.jpg",
	//   Phone: 123456,
   //})

   // R
   //var user models.User
   //v2_db.First(&user, 1) //id

   //v2_db.First(&user, "name=?", "name") //カラムを指定
   //fmt.Println(user)

   // U
   //var user models.User
   //v2_db.First(&user, 1)
   //v2_db.Model(&user).Update("age", 22)
   //v2_db.Model(&user).Update("addr", "zs-xxx")
   //v2_db.Model(&user).Updates(map[string]interface{}{"age":45, "addr": "ssll"})

   // D

   //v2_db.Create(&models.User{
	//   Name:  "削除用",
	//   Age:   11,
	//   Addr:  "ss",
	//   Pic:   "dd",
	//   Phone: 1,
   //})
   //v2_db.Where("name=?", "削除用").Delete(&models.User{})

}
