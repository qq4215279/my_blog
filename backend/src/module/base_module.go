package module

import (
	"gorm.io/gorm"
	"time"
)

type BaseTable struct {
	Id        uint           `gorm:"primarykey" json:"id" form:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
