/**
 * tag
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:40
 */
package module

type Tag struct {
	BaseTable
	// 标签名称
	Name string `json:"name" form:"name"' gorm:"unique_index"`
}

func (tag Tag) TableName() string {
	return "tag"
}
