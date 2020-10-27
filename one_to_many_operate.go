package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_project/models/relate_table"
)

func main(){
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	// C 追加
	//user2 := relate_table.User2{
	//	Model:   gorm.Model{},
	//	Name:    "hallen",
	//	Age:     18,
	//	Addr:    "asaasa",
	//	Article: []relate_table.Article{
	//		{Title:"beego", Content:"beegomorexxx", Desc:"dwwddw"},
	//		{Title:"gin", Content:"ginmoredsds", Desc:"fsdsddsds"},
	//	},
	//}

	//article := relate_table.Article{
	//	Model:   gorm.Model{},
	//	Title:   "タイトル",
	//	Content: "内容",
	//	Desc:    "詳細",
	//}
	//article1 := relate_table.Article{
	//	Model:   gorm.Model{},
	//	Title:   "タイトル1",
	//	Content: "内容1",
	//	Desc:    "詳細1",
	//}
	//
	//user2 := relate_table.User2{
	//	Model:   gorm.Model{},
	//	Name:    "hallen",
	//	Age:     18,
	//	Addr:    "asaasa",
	//	Article: []relate_table.Article{
	//		article,
	//		article1,
	//	},
	//}
	//
	//db.Create(&user2)

	/// R サーチ

	// Preload
	//var user2 relate_table.User2
	//db.Preload("Article").Find(&user2, 1) //SELECT * FROM `user2` WHERE `user2`.`id` = 1 AND `user2`.`deleted_at` IS NULL
	//fmt.Println(user2)
	//t := user2.Article[0].CreatedAt
	//st := t.Format("2006-01-02 15:04:05")
	//fmt.Println(st) // 2020-10-27 14:00:51

	// Association
	//var user3 relate_table.User2
	//db.First(&user3, 1)
	//db.Model(&user3).Association("Article").Find(&user3.Article)
	//fmt.Println(user3)

	//// Updateデータ更新

	//var user3 relate_table.User2
	//db.Preload("Article").Find(&user3, 1)
	//fmt.Println(user3)
	//
	//db.Model(&user3.Article).Where("title = ?", "beego").Update("title", "beego123")

	// D
	var user relate_table.User2
	db.Preload("Article").Find(&user, 1)
	db.Debug().Where("id= ?", "1").Delete(&user.Article)




}
