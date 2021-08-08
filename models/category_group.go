package models

import (
	"github.com/jinzhu/gorm"
)

type CategoryGroup struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;" json:"name"`
	Hidden   bool   `gorm:"not null; default:0;" json:"hidden"`
	Category []Category
}

func (b *CategoryGroup) TableName() string {
	return "category_group"
}
