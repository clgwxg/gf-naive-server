package v1

import (
	"gf-admin/api"
	"gf-admin/internal/model"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseEmp struct {
	Account  string `p:"name" v:"required#请输入账号"`
	NickName string `p:"nickName" v:"required#请输入昵称"`
	DeptId   *int   `p:"deptId" v:"required#请选择部门"`
	RoleId   *int   `P:"roleId"`
	PostId   *int   `p:"postId"`
	Phone    string `p:"phone" v:"required|phone#手机号不能为空|手机号格式错误"`
	Email    string `p:"email" v:"email#邮箱格式错误"`
	Status   uint8  `p:"status"`
	Remark   string `p:"remark"`
}
type CreateReq struct {
	g.Meta   `path:"/emp" tags:"Emp" method:"post" summary:"员工创建" permission:"system:emp:add"`
	Password string `p:"password" v:"required#请输入密码"`
	BaseEmp
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta `path:"/emp" tags:"Emp" method:"put" summary:"员工修改" permission:"system:emp:edit"`
	Id     int `p:"id" v:"required|min:1#数据有误|数据有误"`
	BaseEmp
}
type UpdateRes struct {
}

type GetListReq struct {
	g.Meta `path:"/emp" tags:"Emp" method:"get" summary:"员工列表" permission:"system:emp:list"`
	api.Page
	KeyWord string `p:"keyword"`
	Status  *int8  `p:"status"`
	Phone   string `p:"phone"`
	DeptId  string `p:"deptId"`
}
type GetListRes struct {
	api.PageResult[model.EmpOut]
}

type GetOneReq struct {
	g.Meta `path:"/emp/{id}" tags:"Emp" method:"get" summary:"员工查询" permission:"system:emp:view"`
}
type GetOneRes struct {
	*entity.Emp
}

type DeleteReq struct {
	g.Meta `path:"/emp/{id}" tags:"Emp" method:"delete" summary:"员工删除" permission:"system:emp:delete"`
}
type DeleteRes struct {
}
