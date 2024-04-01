// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Emp is the golang structure for table emp.
type Emp struct {
	Id        int         `json:"id"        ` //
	Account   string      `json:"account"   ` // 账号
	Password  string      `json:"password"  ` // 密码
	NickName  string      `json:"nickName"  ` // 用户昵称
	DeptId    int         `json:"deptId"    ` // 部门id
	RoleId    int         `json:"roleId"    ` // 用户角色
	PostId    int         `json:"postId"    ` // 用户岗位
	Phone     string      `json:"phone"     ` // 手机号
	Email     string      `json:"email"     ` // 邮箱
	Remark    string      `json:"remark"    ` // 备注
	Status    int         `json:"status"    ` // 用户状态
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
