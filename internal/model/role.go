package model

import (
	"gf-admin/api"
	"gf-admin/internal/model/entity"
)

type BaseRoleInput struct {
	Admin   uint8
	Name    string
	Status  uint8
	Remark  string
	MenuIds string
}

type RoleCreateInput struct {
	BaseRoleInput
}

type RoleUpdateInput struct {
	Id int
	BaseRoleInput
}

type RoleQueryInput struct {
	Name   string
	Status *int8
	Date   []string
	api.Page
}

type RoleOut struct {
	Role  *entity.Role
	Menus []entity.RoleMenu
}
