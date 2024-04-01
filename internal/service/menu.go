// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/model/entity"
)

type (
	IMenu interface {
		Create(ctx context.Context, in model.MenuCreateInput) error
		Query(ctx context.Context, in model.MenuQueryInput) ([]*model.MenuTree, error)
		QueryById(ctx context.Context, id int) (*entity.Menu, error)
		Update(ctx context.Context, in model.MenuUpdateInput) error
		DeleteById(ctx context.Context, id int) error
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
