package v1

import (
	"gf-admin/api"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseRole struct {
	Admin   uint8  `p:"admin"`
	Name    string `p:"name" v:"required#请输入角色名称"`
	Status  uint8  `p:"status"`
	Remark  string `p:"remark"`
	MenuIds string `p:"menuIds" v:"required#请选择菜单"`
}
type CreateReq struct {
	g.Meta `path:"/role" tags:"Role" method:"post" summary:"角色创建" permission:"system:role:add"`
	BaseRole
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta `path:"/role" tags:"Role" method:"put" summary:"角色修改" permission:"system:role:edit"`
	Id     int `p:"id" v:"required|min:1#数据有误|数据有误"`
	BaseRole
}
type UpdateRes struct {
}

type GetListReq struct {
	g.Meta `path:"/role" tags:"Role" method:"get" summary:"角色列表" permission:"system:role:list"`
	api.Page
	Name   string   `p:"name"`
	Status *int8    `p:"status"`
	Date   []string `p:"date"`
}
type GetListRes struct {
	api.PageResult[entity.Role]
}

type GetOneReq struct {
	g.Meta `path:"/role/{id}" tags:"Role" method:"get" summary:"角色查询" permission:"system:role:view"`
}
type GetOneRes struct {
	*entity.Role
	MenuIds []int `json:"menuIds"`
}

type DeleteReq struct {
	g.Meta `path:"/role/{id}" tags:"Role" method:"delete" summary:"角色删除" permission:"system:role:delete"`
}
type DeleteRes struct {
}
