package model

import "gf-admin/internal/model/entity"

type BaseMenuInput struct {
	Name     string
	ParentId int64
	Type     uint8
	Icon     string
	IsFrame  uint8
	IsCache  uint8
	Visible  uint8
	Status   uint8
	Url      string
	Path     string
	Query    string
	Perms    string
}

type MenuCreateInput struct {
	BaseMenuInput
}

type MenuQueryInput struct {
	Name   string
	Status *int8
}

type MenuTree struct {
	entity.Menu
	Children []*MenuTree `json:"children"`
}

type MenuUpdateInput struct {
	Id int
	BaseMenuInput
}
