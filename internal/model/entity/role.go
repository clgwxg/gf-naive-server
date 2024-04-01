// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Admin     int         `json:"admin"     ` // 是否超级管理员：1是0否
	Id        int         `json:"id"        ` // 主键id
	Remark    string      `json:"remark"    ` // 备注
	Status    int         `json:"status"    ` // 角色状态：1正常0停用
	Name      string      `json:"name"      ` // 角色名称
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
