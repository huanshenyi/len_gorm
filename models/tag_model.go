package models

import "gorm.io/gorm"

type Article struct {
   gorm.Model
   Title string `gorm:"not null;unique;size:64"`
   Content string `gorm:"column:a_content"`
   Desc string
   Test string `gorm:"-"`
}
