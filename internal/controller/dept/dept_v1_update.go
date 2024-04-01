package dept

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/dept/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, service.Dept().Update(ctx, model.DeptUpdateInput{
		Id: req.Id,
		BaseDeptInput: model.BaseDeptInput{
			Name:     req.Name,
			ParentId: req.ParentId,
			Leader:   req.Leader,
			Status:   req.Status,
			Phone:    req.Phone,
			Email:    req.Email,
		},
	})
}
