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
	IDept interface {
		Create(ctx context.Context, in model.DeptCreateInput) error
		Query(ctx context.Context, in model.DeptQueryInput) ([]*model.DeptTree, error)
		QueryById(ctx context.Context, id int) (*entity.Dept, error)
		Update(ctx context.Context, in model.DeptUpdateInput) error
		DeleteById(ctx context.Context, id int) error
	}
)

var (
	localDept IDept
)

func Dept() IDept {
	if localDept == nil {
		panic("implement not found for interface IDept, forgot register?")
	}
	return localDept
}

func RegisterDept(i IDept) {
	localDept = i
}
