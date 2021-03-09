package module

// 权限
type Privilege struct {
	BaseTable
	// 展示名称
	Label string `json:"label"`
	// path
	Path string `json:"path"`
	// 功能名
	Name string `json:"name"`
	// 图标
	Icon string `json:"icon"`
	// 是否是叶子节点（0：不是，1：是）
	IsLeaf uint `json:"is_leaf"`
	// 是否禁用(0:未禁用，1：禁用)
	IsForbidden uint `json:"is_forbidden"`
	// 父级节点名
	ParentName string `json:"parent_name"`
	// 父节点path
	ParentPath string `json:"parent_path"`
}

// 表名
func (Privilege) TableName() string {
	return "privilege"
}
