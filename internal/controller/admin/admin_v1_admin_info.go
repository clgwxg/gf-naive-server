package admin

import (
	"context"
	"gf-admin/internal/service"

	"gf-admin/api/admin/v1"
)

func (c *ControllerV1) AdminInfo(ctx context.Context, _ *v1.AdminInfoReq) (res *v1.AdminInfoRes, err error) {
	info, err := service.Admin().GetAdminInfo(ctx)
	return &v1.AdminInfoRes{
		AdminInfo: info,
	}, err
}
