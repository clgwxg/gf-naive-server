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
	IEmp interface {
		Create(ctx context.Context, in model.EmpCreateInput) error
		DeleteById(ctx context.Context, id int) error
		Update(ctx context.Context, in model.EmpUpdateInput) error
		Query(ctx context.Context, in model.EmpQueryInput) ([]model.EmpOut, int, error)
		QueryById(ctx context.Context, id int) (*entity.Emp, error)
	}
)

var (
	localEmp IEmp
)

func Emp() IEmp {
	if localEmp == nil {
		panic("implement not found for interface IEmp, forgot register?")
	}
	return localEmp
}

func RegisterEmp(i IEmp) {
	localEmp = i
}
