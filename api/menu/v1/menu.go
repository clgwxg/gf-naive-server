package v1

import (
	"gf-admin/internal/model"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseMenu struct {
	ParentId int64  `p:"parentId"`
	Type     uint8  `p:"type"`
	Icon     string `p:"icon"`
	Name     string `p:"name" v:"required|length:2,6#请输入菜单名称|菜单名称长度为:{min}到:{max}位"`
	IsFrame  uint8  `p:"isFrame"`
	Url      string `p:"url"`
	Path     string `p:"path"`
	Query    string `p:"query"`
	IsCache  uint8  `p:"isCache"`
	Visible  uint8  `p:"visible"`
	Status   uint8  `p:"status"`
	Perms    string `p:"perms"`
}

// 新增
type CreateReq struct {
	g.Meta `path:"/menu" tags:"Menu" method:"post" summary:"菜单创建" permission:"system:menu:add"`
	BaseMenu
}

type CreateRes struct {
}

// 修改
type UpdateReq struct {
	g.Meta `path:"/menu" tags:"Menu" method:"put" summary:"修改创建" permission:"system:menu:edit"`
	Id     int `p:"id" v:"required|min:1#数据有误|数据有误"`
	BaseMenu
}
type UpdateRes struct {
}

// 列表
type GetListReq struct {
	g.Meta `path:"/menu" tags:"Menu" method:"get" summary:"菜单列表" permission:"system:menu:list"`
	Name   string `p:"name"`
	Status *int8  `p:"status"`
}
type GetListRes struct {
	List []*model.MenuTree `json:"list"`
}

// 查询
type GetOneReq struct {
	g.Meta `path:"/menu/{id}" tags:"Menu" method:"get" summary:"菜单查询" permission:"system:menu:view"`
}
type GetOneRes struct {
	*entity.Menu
}

// 删除
type DeleteReq struct {
	g.Meta `path:"/menu/{id}" tags:"Menu" method:"delete" summary:"菜单删除" permission:"system:menu:delete"`
}
type DeleteRes struct {
}
