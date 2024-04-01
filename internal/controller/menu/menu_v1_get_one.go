package menu

import (
	"context"
	"gf-admin/api/menu/v1"
	"gf-admin/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetOne(ctx context.Context, _ *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	id := g.RequestFromCtx(ctx).GetRouter("id")
	menu, err := service.Menu().QueryById(ctx, id.Int())
	return &v1.GetOneRes{Menu: menu}, err
}
