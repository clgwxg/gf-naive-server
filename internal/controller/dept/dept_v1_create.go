package dept

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/dept/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, service.Dept().Create(ctx, model.DeptCreateInput{
		BaseDeptInput: model.BaseDeptInput{
			Name:     req.Name,
			ParentId: req.ParentId,
			Leader:   req.Leader,
			Phone:    req.Phone,
			Status:   req.Status,
			Email:    req.Email,
		},
	})
}
