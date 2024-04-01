package menu

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/menu/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, service.Menu().Update(ctx, model.MenuUpdateInput{
		Id: req.Id,
		BaseMenuInput: model.BaseMenuInput{
			Name:     req.Name,
			Icon:     req.Icon,
			Type:     req.Type,
			ParentId: req.ParentId,
			Url:      req.Url,
			Perms:    req.Perms,
			Path:     req.Path,
			IsFrame:  req.IsFrame,
			IsCache:  req.IsCache,
			Visible:  req.Visible,
			Status:   req.Status,
			Query:    req.Query,
		},
	})
}
