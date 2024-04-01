package role

import (
	"context"
	"gf-admin/internal/dao"
	"gf-admin/internal/model"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
	"gf-admin/internal/service"
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

type sRole struct {
}

func init() {
	service.RegisterRole(&sRole{})
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) error {
	var role *entity.Role
	_ = dao.Role.Ctx(ctx).Where(do.Role{Name: in.Name}).Scan(&role)
	if role != nil {
		return response.NewError("角色名称已存在", nil)
	}
	err := dao.Role.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := dao.Role.Ctx(ctx).Data(do.Role{
			Admin:  in.Admin,
			Name:   in.Name,
			Remark: in.Remark,
			Status: in.Status,
		}).InsertAndGetId()
		if err != nil {
			return err
		}
		split := gstr.Split(in.MenuIds, ",")
		list := make([]do.RoleMenu, 0)
		for _, value := range split {
			list = append(list, do.RoleMenu{MenuId: value, RoleId: id})
		}
		_, err = dao.RoleMenu.Ctx(ctx).Data(list).Insert()
		return err
	})
	return err
}

func (s *sRole) DeleteById(ctx context.Context, id int) error {
	return dao.Role.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Role.Ctx(ctx).Where(do.Role{Id: id}).Delete()
		if err != nil {
			return err
		}
		_, err = dao.RoleMenu.Ctx(ctx).Where(do.RoleMenu{RoleId: id}).Delete()
		return err
	})
}

func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	var role *entity.Role
	_ = dao.Role.Ctx(ctx).Where(do.Role{Name: in.Name}).Where("id !=?", in.Id).Scan(&role)
	if role != nil {
		return response.NewError("角色名称已存在", nil)
	}
	return dao.Role.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Role.Ctx(ctx).Data(do.Role{
			Admin:  in.Admin,
			Name:   in.Name,
			Status: in.Status,
			Remark: in.Remark,
		}).Where(do.Role{Id: in.Id}).Update()
		if err != nil {
			return err
		}
		_, err = dao.RoleMenu.Ctx(ctx).Where(do.RoleMenu{RoleId: in.Id}).Delete()
		if err != nil {
			return err
		}
		split := gstr.Split(in.MenuIds, ",")
		list := make([]do.RoleMenu, 0)
		for _, value := range split {
			list = append(list, do.RoleMenu{MenuId: value, RoleId: in.Id})
		}
		_, err = dao.RoleMenu.Ctx(ctx).Data(list).Insert()
		return err
	})
}

func (s *sRole) Query(ctx context.Context, in model.RoleQueryInput) (list []entity.Role, total int, err error) {
	m := dao.Role.Ctx(ctx).OmitEmpty()
	if in.Name != "" {
		m = m.Where("name like ?", "%"+in.Name+"%")
	}
	if in.Status != nil {
		m = m.Where(do.Role{Status: *in.Status})
	}
	if in.Date != nil {
		m = m.WhereBetween("created_at", in.Date[0], in.Date[1])
	}
	m = m.OrderDesc("updated_At")
	if in.PageNum > 0 {
		err = m.Page(in.PageNum, in.PageSize).
			ScanAndCount(&list, &total, true)
	} else { // in.pageNum 小于0 查询全部
		err = m.Scan(&list)
	}

	if err != nil {
		return nil, 0, response.NewError("查询失败，稍后重试", nil)
	}
	return list, total, nil
}
func (s *sRole) QueryById(ctx context.Context, id int) (*model.RoleOut, error) {
	var role model.RoleOut
	err := dao.Role.Ctx(ctx).Where(do.Role{
		Id: id,
	}).Scan(&role.Role)
	if err != nil || role.Role == nil {
		return nil, err
	}
	err = dao.RoleMenu.Ctx(ctx).Scan(&role.Menus, "role_id", id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *sRole) QueryRoleMenuByRoleId(ctx context.Context, roleId int) ([]*entity.RoleMenu, error) {
	var roleMenus []*entity.RoleMenu
	err := dao.RoleMenu.Ctx(ctx).Where(do.RoleMenu{RoleId: roleId}).Scan(&roleMenus)
	return roleMenus, err
}
