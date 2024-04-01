package role

import (
	"context"
	"gf-admin/internal/service"
	"github.com/gogf/gf/v2/frame/g"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, _ *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	id := g.RequestFromCtx(ctx).GetRouter("id")
	return nil, service.Role().DeleteById(ctx, id.Int())
}
