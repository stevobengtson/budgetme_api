package models

import (
	"github.com/jinzhu/gorm"
)

type Budget struct {
	gorm.Model
	CategoryID uint `gorm:"not null;" json:"category_id"`
	Category   Category
	Month      uint    `gorm:"not null;" json:"month"`
	Year       uint    `gorm:"not null;" json:"year"`
	Assigned   float32 `gorm:"precision:13;scale:4;not null;default:0.0000;" json:"assigned"`
	Activity   float32 `gorm:"precision:13;scale:4;not null;default:0.0000;" json:"activity"`
	Available  float32 `gorm:"precision:13;scale:4;not null;default:0.0000;" json:"available"`
}

func (b *Budget) TableName() string {
	return "budget"
}
