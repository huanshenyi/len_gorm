package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_project/models/relate_table"
)

func main() {
	dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	// 使用するDBを初期化
    //db.AutoMigrate(&relate_table.Article2{}, relate_table.Tag{})

    //////////////////////////////////////// C増加

    //tag1 := relate_table.Tag{
    //	Name:"tag", Desc:"tagdesc",
    //}
    //tag2 := relate_table.Tag{
	//	Name: "", Desc: "",
	//}
    //db.Create(&tag1)
    //db.Create(&tag2)

    //var tag1  relate_table.Tag
    //db.First(&tag1, 1)
	//
    //article := relate_table.Article2{
	//	Title:   "タイトル1",
	//	Content: "内容1",
	//	Desc:    "内容12345",
	//	Tags:    []relate_table.Tag{
	//		tag1,
	//	},
	//}
	//db.Create(&article)

	// ========================= Rサーチ ===========================

	// 1. Preload
    //var article2 relate_table.Article2
    //db.Preload("Tags").Find(&article2, 1)
	//fmt.Println(article2)

	// 2. Association
	//var article2 relate_table.Article2
	//db.Find(&article2, 1)
	//db.Model(&article2).Association("Tags").Find(&article2.Tags)
	//fmt.Println(article2)

	// ================= UPアップデート ======================
	//var article2 relate_table.Article2
	//db.Preload("Tags").Find(&article2, 1)

	//fmt.Println(article2) //{1 タイトル 内容 内容123 [{1 tag tagdesc} {2 tag2 tag2desc}]}
	//db.Debug().Model(&article2.Tags).Where("name=?", "tag").Update("name", "tag1")
	//db.Debug().Model(&article2.Tags).Where("name=?", "tag1").Updates(relate_table.Tag{Name:"123", Desc:"asd"})


	// =============　削除 ====================
    var artice5 relate_table.Article2
	db.Preload("Tags").Find(&artice5, 1)
	db.Model(&artice5).Association("Tags").Delete(&artice5.Tags)
}
