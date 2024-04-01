// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"gf-admin/api/admin/v1"
)

type IAdminV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	AdminInfo(ctx context.Context, req *v1.AdminInfoReq) (res *v1.AdminInfoRes, err error)
	AdminLogout(ctx context.Context, req *v1.AdminLogoutReq) (res *v1.AdminLogoutRes, err error)
}
