package menu

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

type sMenu struct {
}

func init() {
	service.RegisterMenu(&sMenu{})
}

func (s *sMenu) Create(ctx context.Context, in model.MenuCreateInput) error {
	if in.Type != 3 {
		var menu *entity.Menu
		_ = dao.Menu.Ctx(ctx).Where(do.Menu{Name: in.Name}).Scan(&menu)
		if menu != nil {
			return response.NewError("菜单名称已存在", nil)
		}
	}

	if (in.Type == 2) && in.Url == "" {
		return response.NewError("路由不能为空", nil)
	}
	_, err := dao.Menu.Ctx(ctx).Insert(do.Menu{
		Name:     in.Name,
		ParentId: in.ParentId,
		Type:     in.Type,
		Status:   in.Status,
		Visible:  in.Visible,
		IsCache:  in.IsCache,
		IsFrame:  in.IsFrame,
		Url:      in.Url,
		Path:     in.Path,
		Icon:     in.Icon,
		Query:    in.Query,
		Perms:    in.Perms,
	})
	return err
}

func (s *sMenu) Query(ctx context.Context, in model.MenuQueryInput) ([]*model.MenuTree, error) {
	var list []*entity.Menu
	m := dao.Menu.Ctx(ctx)
	if in.Name != "" {
		m = m.Where("name like ?", "%"+in.Name+"%")
	}
	if in.Status != nil {
		m = m.Where(&do.Menu{Status: *in.Status})
	}
	err := m.Order("parent_id asc").Scan(&list)
	if err != nil {
		return nil, err
	}
	return s.getMenuTree(list), nil
}

func (s *sMenu) getMenuTree(list []*entity.Menu) []*model.MenuTree {
	inArr := make([]int, 0)
	var deepFind func(list []*entity.Menu, parentId *int) []*model.MenuTree
	deepFind = func(list []*entity.Menu, parentId *int) []*model.MenuTree {
		var newList []*model.MenuTree
		for _, item := range list {
			if tools.Contain(inArr, item.Id) == false {
				dept := &model.MenuTree{
					Menu: *item,
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

func (s *sMenu) QueryById(ctx context.Context, id int) (*entity.Menu, error) {
	var menu *entity.Menu
	err := dao.Menu.Ctx(ctx).Where(do.Menu{Id: id}).Scan(&menu)
	if err != nil {
		return nil, err
	}
	if menu == nil {
		return nil, response.NewError("参数有误", nil)
	}
	return menu, nil
}

func (s *sMenu) Update(ctx context.Context, in model.MenuUpdateInput) error {
	if in.Type != 3 {
		var menu *entity.Menu
		_ = dao.Menu.Ctx(ctx).Where(do.Menu{
			Name: in.Name,
		}).Scan(&menu)
		if menu != nil && menu.Id != in.Id {
			return response.NewError("菜单名称已存在", nil)
		}
	}
	_, err := dao.Menu.Ctx(ctx).WherePri(in.Id).Data(&do.Menu{
		Name:     in.Name,
		ParentId: in.ParentId,
		Type:     in.Type,
		Icon:     in.Icon,
		Path:     in.Path,
		Url:      in.Url,
		IsCache:  in.IsCache,
		IsFrame:  in.IsFrame,
		Query:    in.Query,
		Status:   in.Status,
		Visible:  in.Visible,
		Perms:    in.Perms,
	}).Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *sMenu) DeleteById(ctx context.Context, id int) error {
	count, _ := dao.Menu.Ctx(ctx).Where(do.Menu{
		ParentId: id,
	}).Count()
	if count > 0 {
		return response.NewError("有子菜单不能删除", nil)
	}
	_, err := dao.Menu.Ctx(ctx).WherePri(id).Delete()
	return err
}
