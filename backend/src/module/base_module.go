/**
 * constants
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:56
 */
package module

import (
	"gorm.io/gorm"
	"time"
)

type BaseTable struct {
	Id uint `gorm:"primary_key" json:"id" form:"id"`
	// 创建时间
	CreatedAt time.Time `json:"createdAt"`
	// 更新时间
	UpdatedAt time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
