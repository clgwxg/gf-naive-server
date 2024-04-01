package dept

import (
	"context"
	"gf-admin/internal/dao"
	"gf-admin/internal/model"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
	"gf-admin/internal/service"
	"gf-admin/utility/response"
	"gf-admin/utility/tools"
)

type sDept struct {
}

func init() {
	service.RegisterDept(&sDept{})
}

func (s *sDept) Create(ctx context.Context, in model.DeptCreateInput) error {
	var menu *entity.Dept
	_ = dao.Dept.Ctx(ctx).Where(do.Dept{Name: in.Name, ParentId: in.ParentId}).Scan(&menu)
	if menu != nil {
		return response.NewError("同级部门名称已存在", nil)
	}
	_, err := dao.Dept.Ctx(ctx).Insert(do.Dept{
		Name:     in.Name,
		ParentId: in.ParentId,
		Status:   in.Status,
		Leader:   in.Leader,
		Phone:    in.Phone,
		Email:    in.Email,
	})
	return err
}

func (s *sDept) Query(ctx context.Context, in model.DeptQueryInput) ([]*model.DeptTree, error) {
	var list []*entity.Dept
	m := dao.Dept.Ctx(ctx)
	if in.Name != "" {
		m = m.Where("name like ?", "%"+in.Name+"%")
	}
	if in.Status != nil {
		m = m.Where(&do.Dept{Status: *in.Status})
	}
	err := m.Order("parent_id asc").Scan(&list)
	if err != nil {
		return nil, err
	}

	return s.getDeptTree(list), nil
}

func (s *sDept) getDeptTree(list []*entity.Dept) []*model.DeptTree {
	inArr := make([]int, 0)
	var deepFind func(list []*entity.Dept, parentId *int) []*model.DeptTree
	deepFind = func(list []*entity.Dept, parentId *int) []*model.DeptTree {
		var newList []*model.DeptTree
		for _, item := range list {
			if tools.Contain(inArr, item.Id) == false {
				dept := &model.DeptTree{
					Dept: *item,
				}
				if parentId == nil || item.ParentId == *parentId {
					inArr = append(inArr, item.Id)
					newList = append(newList, dept)
					children := deepFind(list, &item.Id)
					dept.Children = children
				}
			}

		}
		return newList
	}

	return deepFind(list, nil)
}

func (s *sDept) QueryById(ctx context.Context, id int) (*entity.Dept, error) {
	var menu *entity.Dept
	err := dao.Dept.Ctx(ctx).Where(do.Dept{Id: id}).Scan(&menu)
	if err != nil {
		return nil, err
	}
	if menu == nil {
		return nil, response.NewError("参数有误", nil)
	}
	return menu, nil
}

func (s *sDept) Update(ctx context.Context, in model.DeptUpdateInput) error {
	var menu *entity.Dept
	_ = dao.Dept.Ctx(ctx).Where(do.Dept{
		Name:     in.Name,
		ParentId: in.ParentId,
	}).Scan(&menu)
	if menu != nil && menu.Id != in.Id {
		return response.NewError("部门名称已存在", nil)
	}
	_, err := dao.Dept.Ctx(ctx).WherePri(in.Id).Data(&do.Dept{
		Name:     in.Name,
		ParentId: in.ParentId,
		Status:   in.Status,
		Leader:   in.Leader,
		Phone:    in.Phone,
		Email:    in.Email,
	}).Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *sDept) DeleteById(ctx context.Context, id int) error {
	count, _ := dao.Dept.Ctx(ctx).Where(do.Dept{
		ParentId: id,
	}).Count()
	if count > 0 {
		return response.NewError("有子部门不能删除", nil)
	}
	_, err := dao.Dept.Ctx(ctx).WherePri(id).Delete()
	return err
}
