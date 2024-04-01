package menu

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/menu/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	list, err := service.Menu().Query(ctx, model.MenuQueryInput{
		Name:   req.Name,
		Status: req.Status,
	})
	return &v1.GetListRes{List: list}, err
}
