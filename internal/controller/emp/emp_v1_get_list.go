package emp

import (
	"context"
	"gf-admin/api"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/emp/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	query, i, err := service.Emp().Query(ctx, model.EmpQueryInput{
		KeyWord: req.KeyWord,
		Status:  req.Status,
		Phone:   req.Phone,
		Page:    req.Page,
		DeptId:  req.DeptId,
	})

	return &v1.GetListRes{PageResult: api.PageResult[model.EmpOut]{Total: i, List: query}}, err
}
