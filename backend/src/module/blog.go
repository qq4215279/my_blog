/**
 * blog
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:27
 */
package module

type Blog struct {
	BaseTable
	// 标题
	Title string `json:"title"`
	// 摘要
	Summary string `json:"summary"`
	// 内容
	Contend string `json:"contend"`
	// 分类id
	CategoryId int `json:"categoryId"`
	// 标签ids
	TagIds string `json:"tagIds"`
	// 点击量
	Hits int `json:"hits"`
	// 图片url
	ImageUrl string `json:"imageUrl"`
}

func (blog Blog) TableName() string {
	return "blog"
}
