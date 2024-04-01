package emp

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/service"

	"gf-admin/api/emp/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, service.Emp().Update(ctx, model.EmpUpdateInput{
		Id: req.Id,
		BaseEmpInput: model.BaseEmpInput{
			Account:  req.Account,
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
