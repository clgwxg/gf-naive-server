package post

import (
	"context"
	"gf-admin/api"
	"gf-admin/api/post/v1"
	"gf-admin/internal/model"
	"gf-admin/internal/model/entity"
	"gf-admin/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	list, total, err := service.Post().Query(ctx, model.PostQueryInput{
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
		Page: api.Page{
			PageSize: req.PageSize,
			PageNum:  req.PageNum,
		},
	})
	return &v1.GetListRes{PageResult: api.PageResult[entity.Post]{List: list, Total: total}}, err
}
