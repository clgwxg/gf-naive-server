// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure of table dept for DAO operations like Where/Data.
type Dept struct {
	g.Meta    `orm:"table:dept, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 部门名称
	ParentId  interface{} // 上级部门id
	Leader    interface{} // 部门负责人
	Phone     interface{} // 联系电话
	Email     interface{} // 邮箱
	Status    interface{} // 部门状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
