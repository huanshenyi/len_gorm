package main

import (
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
	"gorm_project/models/relate_table"

	//"gorm_project/models/relate_table"
	"gorm.io/driver/mysql"
	//v1_gorm "github.com/jinzhu/gorm"
)

//type User struct {
//	Id int
//	Name string
//	Age int
//	Addr string
//	Pic string
//	Phone int
//}

func main(){
	// v1のデータベース接続
	//username:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local
	//v1_db, err := v1_gorm.Open("mysql", "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//defer v1_db.Close()
	//v1_db.AutoMigrate(&relate_table.User{}, &relate_table.UserProfile{})
	//v1_db.CreateTable(&User{})
	//v1_db.DropTable("users")
	//b := v1_db.HasTable(&User{})


	// v2のデータベース接続
	dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	v2_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//db, err := v2_db.DB()
	//defer db.Close()
	//v2_db.AutoMigrate(&relate_table.User{}, &relate_table.UserProfile{})
	v2_db.AutoMigrate(&relate_table.User2{}, &relate_table.Article{})
	//v2_db.AutoMigrate(&relate_table.Article2{}, &relate_table.Tag{})
	// CreateTable
    //v2_db.AutoMigrate(&User{})

	// 名前指定のtable追加、v1と互換性あり
	//v2_db.Table("user").Create(&User{})

	// テーブルの削除
    //v2_db.Migrator().DropTable(&User{})

	//b := v2_db.Migrator().HasTable(&User{})
	//fmt.Println(b)


	//v2_db.AutoMigrate(&User{})
}
