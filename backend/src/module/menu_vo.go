/**
 * menu_vo
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/3/9 21:34
 */
package module

// 路由菜单
type Menu struct {
	// 名称
	Label string `json:"label"`
	// 路径
	Path string `json:"path"`
	// 名称
	Name string `json:"name"`
	// 图标
	Icon string `json:"icon"`
	// 子路由
	Children []Menu `json:"children"`
}
