/**
 * category
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:37
 */
package module

type Category struct {
	BaseTable
	// 名称
	Name string `json:"name" form:"name"`
	// 描述
	Describe string `json:"describe" form:"describe"`
	// 文章数量
	Count int `json:"count" form:"count"`
}

func (category Category) TableName() string {
	return "category"
}
