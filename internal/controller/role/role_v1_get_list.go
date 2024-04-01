package role

import (
	"context"
	"gf-admin/api"
	"gf-admin/internal/model"
	"gf-admin/internal/model/entity"
	"gf-admin/internal/service"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	query, i, err := service.Role().Query(ctx, model.RoleQueryInput{
		Name:   req.Name,
		Status: req.Status,
		Date:   req.Date,
		Page:   req.Page,
	})

	return &v1.GetListRes{PageResult: api.PageResult[entity.Role]{Total: i, List: query}}, err
}
