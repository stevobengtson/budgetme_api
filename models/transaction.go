package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	Date       *time.Time `gorm:"not null;" json:"date"`
	CategoryID uint       `gorm:"not null;" json:"category_id"`
	Category   Category
	Outflow    float32 `gorm:"precision:13;scale:4;not null;default:0.0000;" json:"outflow"`
	Inflow     float32 `gorm:"precision:13;scale:4;not null;default:0.0000;" json:"inflow"`
	Payee      uint    `gorm:"not null;" json:"payee"`
	Memo       string  `gorm:"size:1024" json:"memo"`
	Cleared    bool    `json:"cleared"`
}

func (b *Transaction) TableName() string {
	return "transaction"
}
