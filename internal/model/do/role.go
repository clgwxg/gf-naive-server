// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:role, do:true"`
	Admin     interface{} // 是否超级管理员：1是0否
	Id        interface{} // 主键id
	Remark    interface{} // 备注
	Status    interface{} // 角色状态：1正常0停用
	Name      interface{} // 角色名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time //
}
