package post

import (
	"context"
	"gf-admin/internal/service"
	"github.com/gogf/gf/v2/frame/g"

	"gf-admin/api/post/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, _ *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	id := g.RequestFromCtx(ctx).GetRouter("id")
	post, err := service.Post().GetPostById(ctx, id.Int())
	return &v1.GetOneRes{Post: post}, err
}
