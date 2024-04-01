package admin

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/admin/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := service.Admin().Login(ctx, model.LoginInput{
		Account:  req.Account,
		Password: req.Password,
	})

	return &v1.LoginRes{token}, err
}
