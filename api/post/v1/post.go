package v1

import (
	"gf-admin/api"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type BasePost struct {
	Name   string `p:"name" v:"required#请输入岗位名称"`
	Code   string `p:"code" v:"required#请输入岗位编码"`
	Status uint8  `p:"status"`
	Remark string `p:"remark" v:"max-length:250#最长250个字符"`
}

type CreateReq struct {
	g.Meta `path:"/post" tags:"Post" method:"post" summary:"岗位创建" permission:"system:post:add"`
	BasePost
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta `path:"/post" tags:"Post" method:"put" summary:"岗位修改" permission:"system:post:edit"`
	Id     int `p:"id" v:"required|min:1#数据有误|数据有误"`
	BasePost
}
type UpdateRes struct {
}

type GetListReq struct {
	g.Meta `path:"/post" tags:"Post" method:"get" summary:"岗位列表" permission:"system:post:list"`
	api.Page
	Name   string `p:"name"`
	Status *uint8 `p:"status"`
	Code   string `p:"code"`
}
type GetListRes struct {
	api.PageResult[entity.Post]
}

type GetOneReq struct {
	g.Meta `path:"/post/{id}" tags:"Post" method:"get" summary:"岗位查询" permission:"system:post:view"`
}
type GetOneRes struct {
	*entity.Post
}

type DeleteReq struct {
	g.Meta `path:"/post/{id}" tags:"Post" method:"delete" summary:"岗位删除" permission:"system:post:delete"`
}
type DeleteRes struct {
}
