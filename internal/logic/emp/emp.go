package emp

import (
	"context"
	"gf-admin/internal/dao"
	"gf-admin/internal/model"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
	"gf-admin/internal/service"
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

type sEmp struct {
}

func init() {
	service.RegisterEmp(&sEmp{})
}

func (s *sEmp) Create(ctx context.Context, in model.EmpCreateInput) error {
	var role *entity.Emp
	_ = dao.Emp.Ctx(ctx).Where(do.Emp{Account: in.Account}).Scan(&role)
	if role != nil {
		return response.NewError("员工账号已存在", nil)
	}
	password, _ := gmd5.Encrypt(in.Password)
	err := dao.Emp.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Emp.Ctx(ctx).Data(do.Emp{
			Account:  in.Account,
			Password: password,
			NickName: in.NickName,
			DeptId:   in.DeptId,
			PostId:   in.PostId,
			RoleId:   in.RoleId,
			Phone:    in.Phone,
			Email:    in.Email,
			Remark:   in.Remark,
			Status:   in.Status,
		}).Insert()
		return err
	})
	return err
}

func (s *sEmp) DeleteById(ctx context.Context, id int) error {
	return dao.Emp.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Emp.Ctx(ctx).Where(do.Emp{Id: id}).Delete()
		return err
	})
}

func (s *sEmp) Update(ctx context.Context, in model.EmpUpdateInput) error {
	var emp *entity.Emp
	_ = dao.Emp.Ctx(ctx).Where(do.Emp{Account: in.Account}).Where("id !=?", in.Id).Scan(&emp)
	if emp != nil {
		return response.NewError("员工账号已存在", nil)
	}
	return dao.Emp.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Emp.Ctx(ctx).Data(do.Emp{
			Account:  in.Account,
			NickName: in.NickName,
			DeptId:   in.DeptId,
			PostId:   in.PostId,
			RoleId:   in.RoleId,
			Phone:    in.Phone,
			Email:    in.Email,
			Remark:   in.Remark,
			Status:   in.Status,
		}).Where(do.Emp{Id: in.Id}).Update()
		return err
	})
}

func (s *sEmp) Query(ctx context.Context, in model.EmpQueryInput) ([]model.EmpOut, int, error) {
	m := dao.Emp.Ctx(ctx).OmitEmpty()
	if in.KeyWord != "" {
		m = m.WhereLike("account", "%"+in.KeyWord+"%").WhereOrLike("nick_name", "%"+in.KeyWord+"%")
	}
	if in.Phone != "" {
		m = m.WhereLike("phone", "%"+in.Phone+"%")
	}
	if in.Status != nil {
		m = m.Where(do.Emp{Status: &in.Status})
	}
	if in.DeptId != "" {
		m = m.WhereIn("dept_id", gstr.Split(in.DeptId, ","))
	}
	m = m.OrderDesc("created_At").FieldsEx("password")
	total := 0
	var list []model.EmpOut
	err := m.WithAll().Page(in.PageNum, in.PageSize).
		ScanAndCount(&list, &total, true)
	if err != nil {
		return nil, 0, response.NewError("查询失败，稍后重试", nil)
	}
	return list, total, nil
}
func (s *sEmp) QueryById(ctx context.Context, id int) (*entity.Emp, error) {
	var emp *entity.Emp
	err := dao.Emp.Ctx(ctx).Where(do.Emp{
		Id: id,
	}).FieldsEx("password").Scan(&emp)
	return emp, err
}
