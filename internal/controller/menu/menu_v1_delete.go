package menu

import (
	"context"
	"gf-admin/api/menu/v1"
	"gf-admin/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Delete(ctx context.Context, _ *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	id := g.RequestFromCtx(ctx).GetRouter("id")
	return nil, service.Menu().DeleteById(ctx, id.Int())
}
