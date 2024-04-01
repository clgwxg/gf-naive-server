package v1

import (
	"gf-admin/internal/model"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseDept struct {
	ParentId int64  `p:"parentId"`
	Name     string `p:"name" v:"required|length:2,20#请输入部门名称|部门名称长度为:{min}到:{max}位"`
	Status   uint8  `p:"status"`
	Phone    string `p:"phone" v:"phone#手机号格式错误"`
	Leader   string `p:"leader"`
	Email    string `p:"email" v:"email#邮箱格式错误"`
}

// 新增
type CreateReq struct {
	g.Meta `path:"/dept" tags:"Dept" method:"post" summary:"部门创建" permission:"system:dept:add"`
	BaseDept
}

type CreateRes struct {
}

// 修改
type UpdateReq struct {
	g.Meta `path:"/dept" tags:"Dept" method:"put" summary:"部门创建" permission:"system:dept:edit"`
	Id     int `p:"id" v:"required|min:1#数据有误|数据有误"`
	BaseDept
}
type UpdateRes struct {
}

// 列表
type GetListReq struct {
	g.Meta `path:"/dept" tags:"Dept" method:"get" summary:"部门列表" permission:"system:dept:list"`
	Name   string `p:"name"`
	Status *int8  `p:"status"`
}
type GetListRes struct {
	List []*model.DeptTree `json:"list"`
}

// 查询
type GetOneReq struct {
	g.Meta `path:"/dept/{id}" tags:"Dept" method:"get" summary:"部门查询" permission:"system:dept:view"`
}
type GetOneRes struct {
	*entity.Dept
}

// 删除
type DeleteReq struct {
	g.Meta `path:"/dept/{id}" tags:"Dept" method:"delete" summary:"部门删除" permission:"system:dept:delete"`
}
type DeleteRes struct {
}
