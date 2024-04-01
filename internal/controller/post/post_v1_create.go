package post

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/post/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, service.Post().Create(ctx, model.PostCreateInput{
		BasePostInput: model.BasePostInput{
			Name:   req.Name,
			Code:   req.Code,
			Status: req.Status,
			Remark: req.Remark,
		},
	})
}
