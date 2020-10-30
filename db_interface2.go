package main

import (
	"fmt"
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

	//============= where ================
	//var user relate_table.User
	//db.Where("name=?", "").Find(&user)
	// Where("name = ? AND age >= ?", "jinzhu", "22")
	// Where("updated_at > ?", lastWeek)
	// Where("created_at BETWEEN ? AND ?", lastWeek, today)

	// ============= select(指定したカラムだけを取得)  ===================
	//var user relate_table.User
	//db.Select("name,age").Find(&user)
	//fmt.Println(user)

	// create
	//user3 := relate_table.User{
	//	Name:        "hallen8",
	//	Age:         19,
	//	Addr:        "xxx",
	//}
	//db.Create(&user3)
	//user_list := []relate_table.User{
	//	{Name:"1", Age:1, Addr:"1"},
	//	{Name:"2", Age:2, Addr:"2"},
	//}
	//db.Debug().Create(&user_list) //INSERT INTO `users` (`name`,`age`,`addr`) VALUES ('1',1,'1'),('2',2,'2')

	//============= save ===============
	//user4 := relate_table.User{
	//	Name:        "sa",
	//	Age:         11,
	//	Addr:        "sa",
	//}
	//db.Save(&user4)

	//var user relate_table.User
	//db.Where("name = ?", "sa").First(&user)
	//fmt.Println(user)
	//user.Name = "sasa"
	//db.Save(&user) //UPDATE `users` SET `name`='sasa',`age`=11,`addr`='sa' WHERE `id` = 8

	//================== update =====================

	//var user relate_table.User
	//db.Where("age BETWEEN ? and ?", 10, 15).First(&user)
	////db.Where("name=?", "1").First(&user)
	//db.Model(&user).Update("name", "アップデート")

	// 一行もできる
	//db.Where("name=?", "アップデート").Find(&user).Update("name", "update")
	//db.Where("name=?", "アップデート").Find(&user).Updates(map[string]interface{}{"name": "dsa"})

	//================ Delete =======================
	//var user relate_table.User
	//db.Where("name=?", "2").Find(&user).Delete(&user)
	//fmt.Println(user)

	// ============= not ===============
	//var user relate_table.User
	//db.Debug().Not("name=?", "1").Find(&user) //SELECT * FROM `users` WHERE NOT name='1'
	//fmt.Println(user) // [{1 update 10 xx {0    0}} {2 usernmae2 10 xxx {0    0}} {3 name2 19 住所 {0    0}} {4 name21 12 住所1 {0    0}} {5 hallen8 19 xxx {0    0}}]

	//var user1 []relate_table.User
	//db.Debug().Not("name=?", "1").Find(&user1)
	//fmt.Println(user1)

	// ========== or ===========
	//var user []relate_table.User
	//db.Where("name=?", 1).Or("name=?", "name21").Find(&user)
	//db.Where("name = ?", "1").Or("name = ?", "")
    //fmt.Println(user)

    // =========  Order並び替え ===========
    //var user []relate_table.User
	//db.Debug().Where("name LIKE ?", "na%").Order("age desc").Find(&user)
	//fmt.Println(user)

	// ========  limit offset  ========
	//var user []relate_table.User
	//db.Debug().Limit(3).Find(&user) // SELECT * FROM `users` LIMIT 3
	//fmt.Println(user)
	//db.Debug().Limit(2).Offset(2).Find(&user) //SELECT * FROM `users` LIMIT 2 OFFSET 2
	//fmt.Println(user)

	// ========= scan他の構造体にデータを移す ==========
	//var user relate_table.User
	//type User_Scan struct {
	//	Id int
	//	Name string
	//	Age int
	//	Addr string
	//}
	//var userscan User_Scan
	//db.Find(&user).Scan(&userscan)
	//fmt.Println(user, userscan)

	// ======  count =======
	//var user14 []relate_table.User
	//var count int64
	//db.Debug().Model(&user14).Where("name like ?", "name%").Count(&count) //SELECT count(1) FROM `users` WHERE name like 'name%'
	//fmt.Println(count)

	// ======= group having ========
	//var user []relate_table.User
	//type User_Scan struct {
	//	Name string
	//	Age int
	//	Count int
	//}
	//var userscan []User_Scan
	//// SELECT name,age,count(*) as count FROM `users` GROUP BY `name` HAVING age BETWEEN 9 and 15
	//db.Debug().Select("name,age,count(*) as count").Group("name").Having("age BETWEEN ? and ?", 9, 15).Find(&user).Scan(&userscan)
	//fmt.Println(user)
	//fmt.Println(userscan) //[{name21 12 1} {update 10 1} {usernmae2 10 1}]

	// =======  Distinct被るデータを排除 =========
	//var user []relate_table.User
	// SELECT DISTINCT `name`,`age` FROM `users` ORDER BY name, age
    //db.Debug().Distinct("name", "age").Order("name, age").Find(&user)
	//fmt.Println(user) //[{0 1 1  {0    0}} {0 name2 19  {0    0}} {0 name21 12  {0    0}} {0 update 10  {0    0}} {0 usernmae2 10  {0    0}}]

	// ======= join ========
	type UserProfile struct {
		Id int
		Name string
		Age int
		Addr string
		Pic string
		CPic string
		Phone string
	}
	var userprofile []UserProfile
	var user []relate_table.User
	// SELECT * FROM `users` left join user_profiles on users.id = user_profiles.user_id
	db.Debug().Select("users.*, user_profiles.*").Joins("left join user_profiles on users.id = user_profiles.user_id").Find(&user).Scan(&userprofile)
	fmt.Println(user)
	fmt.Println(userprofile) // [{2 usernmae2 10 xxx 1.png 2.png 1213132141} {4 update 10 xx add.png add.png dwwd} {0 name2 19 住所   } {0 name21 12 住所1   } {0 name2 19 xxx   } {0 1 1 1   }]


}

