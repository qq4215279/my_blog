package module

import (
	"gorm.io/gorm"
	"time"
)

type BaseTable struct {
	Id         uint           `gorm:"primarykey" json:"id" form:"id"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `gorm:"index" json:"deleteTime"`
}
