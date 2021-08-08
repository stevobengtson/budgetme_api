package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	CategoryGroupID uint `gorm:"not null;" json:"category_group_id"`
	CategoryGroup   CategoryGroup
	Name            string `gorm:"size:255;not null;" json:"name"`
	Hidden          bool   `gorm:"not null; default:0;" json:"hidden"`
}

func (b *Category) TableName() string {
	return "category"
}
