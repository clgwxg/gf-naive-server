package role

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, service.Role().Create(ctx, model.RoleCreateInput{
		model.BaseRoleInput{
			Admin:   req.Admin,
			Name:    req.Name,
			Status:  req.Status,
			Remark:  req.Remark,
			MenuIds: req.MenuIds,
		},
	})
}
