package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_project/models/relate_table"
)

func main()  {
	dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	// ====== FirstOrInit ======
	//var user relate_table.User
	//db.FirstOrInit(&user, relate_table.User{Name:"name2"}) // 下のと類似
	//db.Where("name=?", "name2").Find(&user)

	// もし結果がなければ、新しく生成される構造体をuserに入れる
	// SELECT * FROM `users` WHERE `users`.`name` = 'name2' AND `users`.`age` = 1 ORDER BY `users`.`id` LIMIT 1
	//db.Debug().FirstOrInit(&user, relate_table.User{Name:"name2", Age:1})
	//fmt.Println(user)


	//=========== attrs =============
	// もし結果がなければ、新しく生成される構造体をuserに入れる
	//var user relate_table.User
	//db.Where("name = ?", "13").Attrs(relate_table.User{Name:"test"}).FirstOrInit(&user)
	//fmt.Println(user)

    // ============ Assign ==============
    // Assign: 見つかったかどうか関係なく、戻り値にAssignの指定パラメータを入れる、あれば置き換え、なければ追加
    //var user relate_table.User
	//db.Where("name = ?", 11).Assign(relate_table.User{Id:100}).FirstOrInit(&user)
	//fmt.Println(user)

	// =============== pluck ================
	//var ages []int
	//var users []relate_table.User
    //db.Model(&users).Where("age < ?", 15).Pluck("age", &ages)
	//fmt.Println(ages) //[10 10 12 1]

	//==============  Scopes ==================
	var users []relate_table.User
	db.Debug().Scopes(GetStatusOk).Find(&users) //SELECT * FROM `users` WHERE age = 10
	fmt.Println(users)

}

func GetStatusOk(db *gorm.DB) *gorm.DB {
	return db.Where("age = ?", 10)
}