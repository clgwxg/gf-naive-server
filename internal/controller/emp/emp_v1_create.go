package emp

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/emp/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, service.Emp().Create(ctx, model.EmpCreateInput{
		BaseEmpInput: model.BaseEmpInput{
			Account:  req.Account,
			Password: req.Password,
			NickName: req.NickName,
			DeptId:   req.DeptId,
			PostId:   req.PostId,
			RoleId:   req.RoleId,
			Phone:    req.Phone,
			Email:    req.Email,
			Status:   req.Status,
			Remark:   req.Remark,
		},
	})
}
