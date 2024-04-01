package model

import (
	"gf-admin/api"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/util/gmeta"
)

type BaseEmpInput struct {
	Account  string
	Password string
	NickName string
	DeptId   *int
	RoleId   *int
	PostId   *int
	Phone    string
	Email    string
	Status   uint8
	Remark   string
}

type EmpCreateInput struct {
	BaseEmpInput
}

type EmpUpdateInput struct {
	Id int
	BaseEmpInput
}

type EmpQueryInput struct {
	KeyWord string
	Status  *int8
	Phone   string
	DeptId  string
	api.Page
}

type EmpOut struct {
	gmeta.Meta `orm:"table:emp"`
	entity.Emp
	Post *entity.Post `json:"post" orm:"with:id=post_id"`
	Role *entity.Role `json:"role" orm:"with:id=role_id"`
}
