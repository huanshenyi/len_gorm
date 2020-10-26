package relate_table

import "gorm.io/gorm"

type Article2 struct {
	gorm.Model
	Tags []Tag `gorm:"many2many:Article2_Tags"`
}

type Tag struct {
	gorm.Model
	Name string
	Desc string
}
