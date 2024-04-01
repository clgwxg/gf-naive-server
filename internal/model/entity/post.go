// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Post is the golang structure for table post.
type Post struct {
	Id        int         `json:"id"        ` //
	Name      string      `json:"name"      ` // 岗位名称
	Code      string      `json:"code"      ` // 岗位编码
	Status    int         `json:"status"    ` // 岗位状态
	Remark    string      `json:"remark"    ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
