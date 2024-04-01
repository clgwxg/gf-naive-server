package menu

import (
	"context"
	"gf-admin/api/menu/v1"
	"gf-admin/internal/model"
	"gf-admin/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, service.Menu().Create(ctx, model.MenuCreateInput{
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
