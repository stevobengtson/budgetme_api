package models

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	UserId   uint    `gorm:"not null;" json:"user_id"`
	Nickname string  `gorm:"size:255;not null;" json:"nickname"`
	Balance  float32 `gorm:"precision:13;scale:4;not null;default:0.0000;" json:"balance"`
	Cleared  float32 `gorm:"precision:13;scale:4;not null;default:0.0000;" json:"cleared"`
	Notes    string  `gorm:"size:1024;" json:"notes"`
	Type     uint    `gorm:"not null" json:"type"`
}

func (b *Account) TableName() string {
	return "account"
}
