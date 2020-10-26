package relate_table
//
//import "gorm.io/gorm"
//
//// 一対多
//
//type User2 struct {
//	gorm.Model
//	Name string
//	Age int
//	Addr string
//	Article []Article `gorm:"foreignKey:UId"`
//}
//
//type Article struct {
//	gorm.Model
//	Title string
//	Content string
//	Desc string
//	UId uint
//}