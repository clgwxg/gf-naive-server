package dept

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/dept/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	list, err := service.Dept().Query(ctx, model.DeptQueryInput{
		Name:   req.Name,
		Status: req.Status,
	})
	return &v1.GetListRes{List: list}, err
}
