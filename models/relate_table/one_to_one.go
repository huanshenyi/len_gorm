package relate_table

import "gorm.io/gorm"

// 属する
//type User struct {
//	Id int
//	Name string
//	Age int
//	Addr string
//}
//
//type UserProfile struct {
//	Id int
//	Pic string
//	CPic string
//	Phone string
//    User User `gorm:"foreignKey:UId;references:Id"` //
//    //UserId int
//    UId int
//}

// 含む
type User struct {
	gorm.Model
	UserProfile UserProfile
}

type UserProfile struct {
	gorm.Model
	Number string
	UserID uint
}

