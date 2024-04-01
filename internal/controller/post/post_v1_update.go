package post

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/post/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, service.Post().Update(ctx, model.PostUpdateInput{
		Id: req.Id,
		BasePostInput: model.BasePostInput{
			Name:   req.Name,
			Status: req.Status,
			Remark: req.Remark,
			Code:   req.Code,
		},
	})
}
