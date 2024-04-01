// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Post is the golang structure of table post for DAO operations like Where/Data.
type Post struct {
	g.Meta    `orm:"table:post, do:true"`
	Id        interface{} //
	Name      interface{} // 岗位名称
	Code      interface{} // 岗位编码
	Status    interface{} // 岗位状态
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
