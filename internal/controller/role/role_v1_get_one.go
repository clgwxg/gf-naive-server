package role

import (
	"context"
	"gf-admin/internal/service"
	"github.com/gogf/gf/v2/frame/g"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, _ *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	id := g.RequestFromCtx(ctx).GetRouter("id")
	role, err := service.Role().QueryById(ctx, id.Int())
	if err != nil || role == nil {
		return nil, err
	}
	var menusIds []int
	if role.Menus != nil {
		for _, item := range role.Menus {
			menusIds = append(menusIds, item.MenuId)
		}
	}
	return &v1.GetOneRes{Role: role.Role, MenuIds: menusIds}, err
}
