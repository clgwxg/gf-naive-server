// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Emp is the golang structure of table emp for DAO operations like Where/Data.
type Emp struct {
	g.Meta    `orm:"table:emp, do:true"`
	Id        interface{} //
	Account   interface{} // 账号
	Password  interface{} // 密码
	NickName  interface{} // 用户昵称
	DeptId    interface{} // 部门id
	RoleId    interface{} // 用户角色
	PostId    interface{} // 用户岗位
	Phone     interface{} // 手机号
	Email     interface{} // 邮箱
	Remark    interface{} // 备注
	Status    interface{} // 用户状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
