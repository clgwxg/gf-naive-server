// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta    `orm:"table:menu, do:true"`
	Id        interface{} //
	Name      interface{} // 菜单名称
	ParentId  interface{} // 父级id
	Type      interface{} // 菜单类型：1目录2菜单3按钮
	Icon      interface{} // 菜单图标
	Url       interface{} // 路由地址
	IsFrame   interface{} // 是否外链：1是0否
	Path      interface{} // 组件路径
	Query     interface{} // 路由参数
	IsCache   interface{} // 是否缓存：1缓存0不缓存
	Perms     interface{} // 权限字符
	Visible   interface{} // 是否显示：1显示0不显示
	Status    interface{} // 菜单状态：1正常0禁用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
