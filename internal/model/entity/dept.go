// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure for table dept.
type Dept struct {
	Id        int         `json:"id"        ` // 主键
	Name      string      `json:"name"      ` // 部门名称
	ParentId  int         `json:"parentId"  ` // 上级部门id
	Leader    string      `json:"leader"    ` // 部门负责人
	Phone     string      `json:"phone"     ` // 联系电话
	Email     string      `json:"email"     ` // 邮箱
	Status    int         `json:"status"    ` // 部门状态
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
}
