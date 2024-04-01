package admin

import (
	"context"
	"gf-admin/api/admin/v1"
	"gf-admin/internal/service"
)

func (c *ControllerV1) AdminLogout(ctx context.Context, _ *v1.AdminLogoutReq) (res *v1.AdminLogoutRes, err error) {
	return nil, service.Admin().Logout(ctx)
}
