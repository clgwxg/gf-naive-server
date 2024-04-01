// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id        int         `json:"id"        ` //
	Name      string      `json:"name"      ` // 菜单名称
	ParentId  int         `json:"parentId"  ` // 父级id
	Type      int         `json:"type"      ` // 菜单类型：1目录2菜单3按钮
	Icon      string      `json:"icon"      ` // 菜单图标
	Url       string      `json:"url"       ` // 路由地址
	IsFrame   int         `json:"isFrame"   ` // 是否外链：1是0否
	Path      string      `json:"path"      ` // 组件路径
	Query     string      `json:"query"     ` // 路由参数
	IsCache   int         `json:"isCache"   ` // 是否缓存：1缓存0不缓存
	Perms     string      `json:"perms"     ` // 权限字符
	Visible   int         `json:"visible"   ` // 是否显示：1显示0不显示
	Status    int         `json:"status"    ` // 菜单状态：1正常0禁用
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
}
