package models

import "gorm.io/gorm"

type GormModel struct {
	gorm.Model
	Name string
}

func (GormModel) TableName() string {
	return "test_gorm_model"
}
