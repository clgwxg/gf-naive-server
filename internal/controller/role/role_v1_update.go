package role

import (
	"context"
	"gf-admin/api/role/v1"
	"gf-admin/internal/model"
	"gf-admin/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		BaseRoleInput: model.BaseRoleInput{
			Admin:   req.Admin,
			Name:    req.Name,
			Status:  req.Status,
			MenuIds: req.MenuIds,
			Remark:  req.Remark,
		},
	})
}
