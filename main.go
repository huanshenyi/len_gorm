package main

func main(){
	// v1のデータベース接続
	// username:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local
	//v1_db, err := v1_gorm.Open("mysql", "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//defer v1_db.Close()


	// v2のデータベース接続
	//dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	//v2_db, err := v2_gorm.Open(mysql.Open(dsn), &v2_gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//db, err := v2_db.DB()
	//defer db.Close()
}
