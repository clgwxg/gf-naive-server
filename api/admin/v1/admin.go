package v1

import (
	"gf-admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"login" method:"post" summary:"登录"`
	Account  string `p:"account"  v:"required#请输入账号"`
	Password string `p:"password" v:"required#请输入密码"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type AdminInfoReq struct {
	g.Meta `path:"/info" tags:"userInfo" method:"get" summary:"登录用户信息" auth:"true"`
}

type AdminInfoRes struct {
	*model.AdminInfo
}

type AdminLogoutReq struct {
	g.Meta `path:"/logout" tags:"userInfo" method:"get" summary:"登录用户信息" auth:"true"`
}

type AdminLogoutRes struct {
}
