package relate_table

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


type User struct {
	Id int
	Name string
	Age int
	Addr string
	UserProfile UserProfile
}

type UserProfile struct {
	Id int
	Pic string
	CPic string
	Phone string
	UserId uint
}
