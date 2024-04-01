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
	IRole interface {
		Create(ctx context.Context, in model.RoleCreateInput) error
		DeleteById(ctx context.Context, id int) error
		Update(ctx context.Context, in model.RoleUpdateInput) error
		Query(ctx context.Context, in model.RoleQueryInput) (list []entity.Role, total int, err error)
		QueryById(ctx context.Context, id int) (*model.RoleOut, error)
		QueryRoleMenuByRoleId(ctx context.Context, roleId int) ([]*entity.RoleMenu, error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
