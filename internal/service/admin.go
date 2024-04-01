// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-admin/internal/model"
)

type (
	IAdmin interface {
		Login(ctx context.Context, in model.LoginInput) (token string, err error)
		GetAdminInfo(ctx context.Context) (*model.AdminInfo, error)
		Logout(ctx context.Context) error
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
